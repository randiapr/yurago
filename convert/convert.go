package convert

import (
	"log"
	"strconv"
	"time"
)

const (
	// ISO8601DT date time format
	ISO8601DT string = "2006-01-02 15:04:05"
	// ISO8601D date format
	ISO8601D string = "2006-01-02"
)

// StrToDtISO8601 string to datetime yyyy-MM-dd HH:mm:ss
func StrToDtISO8601(param string) time.Time {
	t, err := time.Parse(ISO8601DT, param)
	if err != nil {
		log.Printf("ERROR: %s", err.Error())
	}
	return t
}

// StrToDateISO8601 string to date yyyy-MM-dd
func StrToDateISO8601(param string) time.Time {
	t, err := time.Parse(ISO8601D, param)
	if err != nil {
		log.Printf("ERROR: %s", err.Error())
	}
	return t
}

// StrToInt64  string to Int64
func StrToInt64(param string) int64 {
	result, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		log.Printf("ERROR: %s", err.Error())
		return 0
	}
	return result
}

// StrToFloat64 string to float64
func StrToFloat64(param string) float64 {
	result, err := strconv.ParseFloat(param, 64)
	if err != nil {
		log.Printf("ERROR: %s", err.Error())
		return 0
	}
	return result
}

// StrToFloat32 string to float32
func StrToFloat32(param string) float32 {
	result, err := strconv.ParseFloat(param, 32)
	if err != nil {
		log.Printf("ERROR: %s", err.Error())
		return 0
	}
	return float32(result)
}

// StrToInt string to int
func StrToInt(param string) int {
	result, err := strconv.Atoi(param)
	if err != nil {
		log.Printf("ERROR: %s", err.Error())
		return 0
	}
	return result
}

// MillisToTime long date to Date time
func MillisToTime(millis int64) time.Time {
	return time.Unix(0, millis*int64(time.Millisecond))
}

// DtToMillis Date time to millisec
func DtToMillis(dt time.Time) int64 {
	return dt.UnixNano() / int64(time.Millisecond)
}

// StrToMillis string date to millis
func StrToMillis(date string) int64 {
	dt, _ := time.Parse(ISO8601DT, date)
	return DtToMillis(dt)
}

func StrDateTimeToMillis(date string) int64 {
	dt, _ := time.Parse(ISO8601D, date)
	return DtToMillis(dt)
}
