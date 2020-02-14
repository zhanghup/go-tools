package tools

import (
	"time"
)

func format(t *time.Time, fmt string) string {
	if t == nil {
		t2 := time.Now()
		t = &t2
	}
	return (*t).Format(fmt)
}

func TimeToHMS(t *time.Time) string {
	return format(t, "15:04:05")
}
func TimeToYMD(t *time.Time) string {
	return format(t, "2006-01-02")
}
func TimeToYM(t *time.Time) string {
	return format(t, "2006-01")
}
func TimeToYear(t *time.Time) string {
	return format(t, "2006")
}
