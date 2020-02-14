package tools

import (
	"time"
)

func TimeToHMS() string {
	return time.Now().Format("15:04:05")
}
func TimeToYMD() string {
	return time.Now().Format("2006-01-02")
}
func TimeToYM() string {
	return time.Now().Format("2006-01")
}
func TimeToYear() string {
	return time.Now().Format("2006")
}
