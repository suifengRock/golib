package main

import (
	"fmt"
	"time"
	"github.com/suifengRock/go-tools"
)

type logicalFunc func()

func Test() {
	fmt.Println("1230.0")
}

func handle(f logicalFunc) {
	go f()
}

func tickerHandle(f logicalFunc, d time.Duration) {
	timer2 := time.NewTicker(d)
	for {
		select {
		case <-timer2.C:
			handle(f)
		}
	}
}

func timerHandle(f logicalFunc, d time.Duration) *time.Timer {
	return time.AfterFunc(d, f)
}

//fmtTime: "15:04:56"
func AlarmClockDaily(f logicalFunc, fmtTime string) {

}

func main() {
	fmt.Println(tools.NoSpeStr(9))

	fmt.Println(tools.GetNowInterval("18:13:00"))
}
