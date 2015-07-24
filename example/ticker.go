package main

import (
	"golib/log"
	"golib/ticker"
	"time"
)

func sayFirst() {
	log.Info("...test...ticker...1")
}

func saySecond() {
	log.Info("...test...ticker...2")
}

func sayThird() {
	log.Info("...test...ticker...3")
}

func sayFour() {
	log.Info("...test...ticker...4")
}

func main() {
	go ticker.TimerHandle(sayFirst, time.Duration(10)*time.Second)

	go ticker.AlarmClockDaily(saySecond, "14:57:00")

	go ticker.AlarmClockonce(sayThird, "14:57:30")

	go ticker.TickerHandle(sayFour, time.Duration(10)*time.Second)

	time.Sleep(time.Duration(2) * time.Minute)
	log.Trace("....end")
}
