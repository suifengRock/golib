package main

import (
	"fmt"
	"golib/time"
)

func main() {

	nowStr := time.FmtTime(0, time.SecondFmt)
	fmt.Println(nowStr)

	nowUnix := time.TimeToUnix(nowStr, time.SecondFmt)
	fmt.Println(nowUnix)

}
