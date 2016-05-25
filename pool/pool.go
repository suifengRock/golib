package services_uCenter

import (
	"container/list"
	"errors"
	"sync"
	"time"
)

type Pool struct {

	// Dial is an application supplied function for creating new connections.
	Dial func() (interface{}, error)

	// Close is an application supplied functoin for closeing connections.
	Close func(c interface{}) error

	// TestOnBorrow is an optional application supplied function for checking
	// the health of an idle connection before the connection is used again by
	// the application. Argument t is the time that the connection was returned
	// to the pool. If the function returns an error, then the connection is
	// closed.
	TestOnBorrow func(c interface{}, t time.Time) error

	// Maximum number of idle connections in the pool.
	MaxIdle int

	// Maximum number of connections allocated by the pool at a given time.
	// When zero, there is no limit on the number of connections in the pool.
	MaxActive int

	// Close connections after remaining idle for this duration. If the value
	// is zero, then idle connections are not closed. Applications should set
	// the timeout to a value less than the server's timeout.
	IdleTimeout time.Duration

	// mu protects fields defined below.
	mu     sync.Mutex
	closed bool
	active int

	// Stack of idleConn with most recently used at the front.
	idle list.List
}

type idleConn struct {
	c interface{}
	t time.Time
}

var (
	ErrPoolExhausted = errors.New("thrift: connection pool exhausted")
	ErrPoolClosed    = errors.New("thrift: connection pool closed")
)
var nowFunc = time.Now // for testing
// New creates a new pool. This function is deprecated. Applications should
// initialize the Pool fields directly as shown in example.
func NewPool(dialFn func() (interface{}, error),
	testFunc func(c interface{}, t time.Time) error,
	closeFn func(c interface{}) error,
	maxIdle int, maxActive int, timeout time.Duration) *Pool {
	return &Pool{
		Dial:         dialFn,
		TestOnBorrow: testFunc,
		Close:        closeFn,
		MaxIdle:      maxIdle,
		MaxActive:    maxActive,
		IdleTimeout:  timeout,
	}
}

// Get gets a connection. The application must close the returned connection.
// This method always returns a valid connection so that applications can defer
// error handling to the first use of the connection.
func (p *Pool) Get() (interface{}, error) {
	p.mu.Lock()

	if p.closed {
		p.mu.Unlock()
		return nil, ErrPoolClosed
	}
	// Prune stale connections.
	if timeout := p.IdleTimeout; timeout > 0 {
		for i, n := 0, p.idle.Len(); i < n; i++ {
			e := p.idle.Back()
			if e == nil {
				break
			}
			ic := e.Value.(idleConn)
			if ic.t.Add(timeout).After(nowFunc()) {
				break
			}
			p.idle.Remove(e)
			p.active -= 1
			p.mu.Unlock()
			p.Close(ic.c)
			p.mu.Lock()
		}
	}
	// Get idle connection.
	for i, n := 0, p.idle.Len(); i < n; i++ {
		e := p.idle.Front()
		if e == nil {
			break
		}
		ic := e.Value.(idleConn)
		p.idle.Remove(e)
		test := p.TestOnBorrow
		p.mu.Unlock()
		if test == nil || test(ic.c, ic.t) == nil {
			return ic.c, nil
		}
		p.Close(ic.c)
		p.mu.Lock()
	}
	// No idle connection, create new.
	if p.MaxActive == 0 || p.active < p.MaxActive {
		dial := p.Dial
		p.active += 1
		p.mu.Unlock()
		c, err := dial()
		if err != nil {
			p.mu.Lock()
			p.active -= 1
			p.mu.Unlock()
			c = nil
		}
		return c, err
	} else {
		p.mu.Unlock()
		return nil, ErrPoolExhausted
	}

}

// Put adds conn back to the pool, use forceClose to close the connection forcely
func (p *Pool) Put(c interface{}, forceClose bool) error {
	if !forceClose && c != nil {
		p.mu.Lock()
		if !p.closed {
			p.idle.PushFront(idleConn{t: nowFunc(), c: c})
			if p.idle.Len() > p.MaxIdle {
				// remove exceed conn
				c = p.idle.Remove(p.idle.Back()).(idleConn).c
			} else {
				c = nil
			}
		}
		p.mu.Unlock()
	}
	// close exceed conn
	if c != nil {
		p.mu.Lock()
		p.active -= 1
		p.mu.Unlock()
		return p.Close(c)
	}
	return nil
}

// ActiveCount returns the number of active connections in the pool.
func (p *Pool) ActiveCount() int {
	p.mu.Lock()
	active := p.active
	p.mu.Unlock()
	return active
}

// Relaase releases the resources used by the pool.
func (p *Pool) Release() error {
	p.mu.Lock()
	idle := p.idle
	p.idle.Init()
	p.closed = true
	p.active -= idle.Len()
	p.mu.Unlock()
	for e := idle.Front(); e != nil; e = e.Next() {
		p.Close(e.Value.(idleConn).c)
	}
	return nil
}
