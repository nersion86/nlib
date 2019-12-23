package timefm

import (
	"time"
)

/* timefm is Time Formatter */

//GetDefaultTimeFormat return YYYY-MM-DD now time
func GetDefaultTimeFormat() string {
	now := time.Now()
	return GetDefaultTimeFormatFromTime(now)
}

//GetDefaultFileTimeFormat return YYYYMMDD now time
func GetDefaultFileTimeFormat() string {
	now := time.Now()
	return GetDefaultFileTimeFormatFromTime(now)
}

//GetDefaultTimeFormatFromTime return YYYY-MM-DD from input time
func GetDefaultTimeFormatFromTime(t time.Time) string {
	return t.Format("2006-01-02")
}

//GetDefaultFileTimeFormatFromTime return YYYYMMDD from input time
func GetDefaultFileTimeFormatFromTime(t time.Time) string {
	return t.Format("20060102")
}

//GetConvertTimeFormat it only provides formats that are often used case.
func GetConvertTimeFormat(t time.Time, str string) string {

	switch str {
	case "YYYY-MM-DD":
		return t.Format("2006-01-02")
	case "YYYY-MM-DD hh":
		return t.Format("2006-01-02 15")
	case "YYYY-MM-DD hh:mm":
		return t.Format("2006-01-02 15:04")
	case "YYYY-MM-DD hh:mm:ss":
		return t.Format("2006-01-02 15:04:05")

	case "YYYYMMDD":
		return t.Format("20060102")
	case "YYYYMMDDhh":
		return t.Format("2006010215")
	case "YYYYMMDDhhmm":
		return t.Format("200601021504")
	case "YYYYMMDDhhmmss":
		return t.Format("20060102150405")

	case "hh":
		return t.Format("15")
	case "hh:mm":
		return t.Format("15:04")
	case "hh:mm:ss":
		return t.Format("15:04:05")
	}

	return ""
}
