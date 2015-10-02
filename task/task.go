package task

import (
	"errors"
	"fmt"
	"time"
)

const (
	NO_START = 1
	EXECUTE  = 2
	STOP     = 3
	FINISH   = 4
)

type logicalFunc func() err

type TaskTicker struct {
	ticker         *time.Ticker  // 可以理解为计时器，闹钟
	function       logicalFunc   // 要执行的func
	afterTimeExec  time.Duration // 多少时间后执行
	intervalTime   time.Duration // 间隔多少时间循环执行 为0：不循环，只执行一次，AllowExecCount失效
	execCount      uint64        // 执行次数
	allowExecCount uint64        // 允许执行次数 为0, 不限次数数循环
	createAt       time.Time     // 创建时间
	stopAt         time.Time     // 停止时间
	finishAt       time.Time     // 完成时间
	execTimeStr    string        // 执行时间
	status         int           // 任务状态
	log            *ExecLog      //执行日志
}

func (t *TaskTicker) GetExecCount() uint64 {
	return t.execCount
}

// *****

func NewTaskTicker(f logicalFunc, after, interval time.Duration, allowCount uint64, log *ExecLog) *TaskTicker {
	if interval == 0 {
		allowCount = 1
	}
	execTimeStr := fmt.Sprintf("%d秒后执行", after/time.Second)
	ticker := new(TaskTicker)
	ticker.afterTimeExec = after
	ticker.intervalTime = interval
	ticker.createAt = time.Now()
	ticker.execTimeStr = execTimeStr
	ticker.allowExecCount = allowCount
	ticker.function = f
	ticker.status = NO_START
	ticker.log = log
	return ticker
}

func (t *TaskTicker) execStart() (err error) {
	if t.status != EXECUTE {
		return errors.New{"the task is not allow execute."}
	}
	if t.allowExecCount <= t.execCount {
		if t.ticker != nil {
			t.ticker.Stop()
		}
		return errors.New(" not allow execute:allow limit.")
	}
	t.execCount += 1
	return
}

func (t *TaskTicker) execLoop() {
	if t.allowExecCount <= t.execCount {
		return
	}
	ticker := time.NewTicker(t.intervalTime)
	for {
		select {
		case <-ticker.C:
			err := t.execStart()
			if err == nil {
				go t.execEnd()
			}
		}
	}
}

func (t *TaskTicker) execEnd() {
	t.log.TickerBegin()
	err = t.function()
	if t.allowExecCount <= t.execCount {
		if t.ticker != nil {
			t.ticker.Stop()
		}
		t.status = FINISH
		t.finishAt = time.Now()
	}
	t.log.TickerEnd()
	return
}

// 计划任务启动
func (t *TaskTicker) taskFire() {
	t.status = EXECUTE
	timer := time.NewTimer(t.afterTimeExec)
	select {
	case <-timer.C:
		err := t.execStart()
		if err == nil {
			go t.execEnd()
			go t.execLoop()
		}
	}
}

// 计划任务停止
func (t *TaskTicker) taskStop() {
	if t.ticker != nil {
		t.ticker.Stop()
	}
	t.status = STOP
	t.stopAt = time.Now()
}
