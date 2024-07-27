package util

import (
	"strings"
	"time"
)

func TimeFormat(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}

func IsInArrayIgnoreCase(a string, arrs []string) bool {
	a = strings.TrimSpace(a)
	a = strings.ToLower(a)
	for _, v := range arrs {
		v = strings.TrimSpace(v)
		if a == strings.ToLower(v) {
			return true
		}
	}
	return false
}
