package util

import (
	"time"
)

func GetNowStr() string {
	return time.Now().Format(time.RFC3339)
}

func ParseStrToTime(t string) (time.Time, error) {
	return time.Parse(time.RFC3339, t)
}
