package ticker

import (
	myTime "golib/time"
	"time"
)

type logicalFunc func()

//  每间隔多少时间执行（循环不断）
func TickerHandle(f logicalFunc, d time.Duration) {
	ticker := time.NewTicker(d)
	for {
		select {
		case <-ticker.C:
			go f()
		}
	}
}

// 多少时间后执行一次
func TimerHandle(f logicalFunc, d time.Duration) *time.Timer {
	return time.AfterFunc(d, f)
}

//  fmtTime: "15:04:05"
//  every day execute the logicalFunc on the fmtTime
func AlarmClockDaily(f logicalFunc, fmtTime string) {
	nowInterval := myTime.GetNowInterval(fmtTime)
	timer := time.NewTimer(time.Duration(nowInterval) * time.Second)
	select {
	case <-timer.C:
		go f()
		go TickerHandle(f, time.Duration(24)*time.Hour)
	}
}

//  fmtTime: "15:04:05"
//  execute the logicalFunc on the fmtTime today, if fmtTime is overdue
//  then execute the logicalFunc on the fmtTime tomorrow.
func AlarmClockonce(f logicalFunc, fmtTime string) *time.Timer {
	nowInterval := myTime.GetNowInterval(fmtTime)
	return TimerHandle(f, time.Duration(nowInterval)*time.Second)
}
