package tools

import "time"

type mytime struct{}

var Ti = mytime{}

func (this mytime) format(t []time.Time, fmt string) string {
	var tt time.Time
	if t != nil && len(t) >= 0 {
		tt = t[0]
	} else {
		tt = time.Now()
	}
	return tt.Format(fmt)
}

func (this mytime) HMS(t ...time.Time) string {
	return this.format(t, "15:04:05")
}
func (this mytime) YMD(t ...time.Time) string {
	return this.format(t, "2006-01-02")
}
func (this mytime) YM(t ...time.Time) string {
	return this.format(t, "2006-01")
}
func (this mytime) Year(t ...time.Time) string {
	return this.format(t, "2006")
}
