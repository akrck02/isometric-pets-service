package utils

import "time"

func GetCurrentMillis() int64 {
	return time.Now().UnixMilli()
}

func CurrentDate() string {
	return time.Now().Format("2006-01-02")
}
