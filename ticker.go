package main

import (
	"fmt"
	"time"

	"github.com/suifengRock/golib"
)

type logicalFunc func()

func Test() {
	fmt.Println("1230.0")
}

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
	nowInterval := golib.GetNowInterval(fmtTime)
	timer := TimerHandle(f, time.Duration(nowInterval)*time.Second)
	select {
	case <-timer.C:
		// TickerHandle(f, time.Duration(24)*time.Hour)
		fmt.Println("...")
	}
}

//fmtTime: "15:04:05"
func AlarmClockonce(f logicalFunc, fmtTime string) *time.Timer {
	nowInterval := golib.GetNowInterval(fmtTime)
	return TimerHandle(f, time.Duration(nowInterval)*time.Second)
}

func main() {
	fmt.Println(golib.NoSpeStr(9))
	AlarmClockDaily(Test, "20:45:56")
	time.Sleep(time.Second * 60)
}
