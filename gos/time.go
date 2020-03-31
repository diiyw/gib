package gos

import "time"

const (
	DateFormat     = "2006-01-02"
	DateTimeFormat = "2006-01-02 15:04:05"
	ZeroTime       = "0000-00-00 00:00:00"
)

// Date return now date
func Date() string {
	return time.Now().Format(DateFormat)
}

// DateTime return string datetime
func DateTime() string {
	return time.Now().Format(DateTimeFormat)
}

func Yesterday() time.Time {
	return time.Now().Add(-time.Hour * 24)
}

func YesterdayDate() string {
	return Yesterday().Format(DateFormat)
}

func ParseTime(layout, v string) time.Time {
	t, _ := time.Parse(layout, v)
	return t
}

func UnixTime(unix int64) time.Time {
	return time.Unix(unix, 0)
}
