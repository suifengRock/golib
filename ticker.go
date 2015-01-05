package golib

import (
	"time"
)

type logicalFunc func()

func handle(f logicalFunc) {
	go f()
}

func TickerHandle(f logicalFunc, d time.Duration) {
	timer := time.NewTicker(d)
	for {
		select {
		case <-timer.C:
			handle(f)
		}
	}
}

func TimerHandle(f logicalFunc, d time.Duration) *time.Timer {
	return time.AfterFunc(d, f)
}

//fmtTime: "15:04:05"
func AlarmClockDaily(f logicalFunc, fmtTime string) {
	nowInterval := GetNowInterval(fmtTime)
	timer := time.NewTimer(time.Duration(nowInterval) * time.Second)
	select {
	case <-timer.C:
		go f()
		go TickerHandle(f, time.Duration(24)*time.Hour)
	}
}

//fmtTime: "15:04:05"
func AlarmClockonce(f logicalFunc, fmtTime string) *time.Timer {
	nowInterval := GetNowInterval(fmtTime)
	return TimerHandle(f, time.Duration(nowInterval)*time.Second)
}
