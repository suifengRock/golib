package golib

import (
	"strconv"
	"time"
)

var (
	YearFmt       = "2006"
	MonFmt        = "2006-01"
	DayFmt        = "2006-01-02"
	HourFmt       = "2006-01-02 15"
	MinuteFmt     = "2006-01-02 15:04"
	SecondFmt     = "2006-01-02 15:04:05"
	OnlySecondFmt = "15:04:05"
	MonthStr      = [13]string{"00", "01", "02", "03", "04", "05", "06",
		"07", "08", "09", "10", "11", "12"}
	MonthDay = [13]int{29, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
)

func FmtTime(stamp int64, format string) string {
	if stamp <= 0 {
		t := time.Now()
		return t.Format(format)
	} else {
		t := time.Unix(int64(stamp), 0)
		return t.Format(format)
	}
}

func TimeToUnix(s string, format string) int64 {
	loc, _ := time.LoadLocation("Local")
	t, err := time.ParseInLocation(format, s, loc)
	if err == nil {
		return t.Unix()
	} else {
		return 0
	}
}

func GetYearAndMonth(date int64) (year, month int) {
	t := time.Unix(int64(date), 0)
	year = t.Year()
	month = int(t.Month())

	return
}

func GetYearUnix(date int64) int64 {

	t := time.Unix(int64(date), 0)
	year := t.Year()

	yearStr := strconv.Itoa(year)
	dateStr := yearStr

	return TimeToUnix(dateStr, YearFmt)

}

func GetMonthUnix(date int64) int64 {

	t := time.Unix(int64(date), 0)
	year := t.Year()
	month := int(t.Month())

	yearStr := strconv.Itoa(year)
	monthStr := MonthStr[month]
	dateStr := yearStr + "-" + monthStr

	return TimeToUnix(dateStr, MonFmt)

}

func GetNextMonthUnix(date int64) int64 {

	t := time.Unix(int64(date), 0)
	t = t.AddDate(0, 1, 0)
	return t.Unix()
}

func GetDayUnix(date int64) int64 {

	dayStr := FmtTime(0, DayFmt)
	return TimeToUnix(dayStr, DayFmt)
}

func GetNowUnix(format string) int64 {
	now := FmtTime(0, format)

	return TimeToUnix(now, format)
}

//fmtTime: "15:04:05"
func GetNowInterval(fmtTime string) int64 {
	dayStr := FmtTime(0, DayFmt)
	timeStr := dayStr + " " + fmtTime
	timeUnix := TimeToUnix(timeStr, SecondFmt)
	nowUnix := GetNowUnix(SecondFmt)
	if timeUnix >= nowUnix {
		return timeUnix - nowUnix
	}

	return timeUnix + 24*60*60 - nowUnix
}

func IsLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

func GetMonthWorkDay(date int64) int {
	t := time.Unix(int64(date), 0)
	firstDay := int(t.Weekday())
	month := int(t.Month())
	year := t.Year()
	monDay := 0
	if month == 2 {
		if IsLeapYear(year) {
			monDay = MonthDay[0]
		}
	} else {
		monDay = MonthDay[month]
	}
	workDay := monDay / 7 * 5
	remainder := month % 7
	if remainder != 0 {
		tmp := remainder + firstDay - 1
		if firstDay == 0 {
			workDay += remainder - 1
		} else if tmp < 6 {
			workDay += remainder
		} else if tmp == 6 {
			workDay += remainder - 1
		} else {
			workDay += remainder - 2
		}
	}
	return workDay
}
