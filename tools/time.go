package tools

import (
	"strconv"
	"time"
)

var (
	YearFmt   = "2006"
	MonFmt    = "2006-01"
	DayFmt    = "2006-01-02"
	HourFmt   = "2006-01-02 15"
	MinuteFmt = "2006-01-02 15:04"
	SecondFmt = "2006-01-02 15:04:56"
	MonthStr  = [13]string{"00", "01", "02", "03", "04", "05", "06",
		"07", "08", "09", "10", "11", "12"}
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

func GetDayUnix(date int64) int64 {

	t := time.Unix(int64(date), 0)
	year := t.Year()
	month := int(t.Month())
	day := t.Day()

	yearStr := strconv.Itoa(year)
	monthStr := MonthStr[month]
	dayStr := strconv.Itoa(day)

	dateStr := yearStr + "-" + monthStr + "-" + dayStr

	return TimeToUnix(dateStr, DayFmt)

}

func GetNowUnix(format string) int64 {
	now := FmtTime(0, format)

	return TimeToUnix(now, format)
}
