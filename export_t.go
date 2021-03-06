package tools

import "time"

type mytime struct{}

var Time = mytime{}

func (this mytime) format(t []time.Time, fmt string) string {
	var tt time.Time
	if t != nil && len(t) >= 0 {
		tt = t[0]
	} else {
		tt = time.Now()
	}
	return tt.Format(fmt)
}

func (this mytime) UnixYMDHMS(t int64) string {
	return this.YMDHMS(time.Unix(t, 0))
}
func (this mytime) UnixHMS(t int64) string {
	return this.HMS(time.Unix(t, 0))
}
func (this mytime) UnixYMD(t int64) string {
	return this.YMD(time.Unix(t, 0))
}
func (this mytime) UnixYM(t int64) string {
	return this.YM(time.Unix(t, 0))
}
func (this mytime) UnixYear(t int64) string {
	return this.Year(time.Unix(t, 0))
}

func (this mytime) YMDHMS(t ...time.Time) string {
	return this.format(t, "2006-01-02 15:04:05")
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
func (this mytime) ParseHMS(i string) (time.Time, error) {
	return this.Parse(i, "15:04:05")
}
func (this mytime) MustParseHMS(i string) time.Time {
	t, _ := this.ParseHMS(i)
	return t
}
func (this mytime) ParseYMD(i string) (time.Time, error) {
	return this.Parse(i, "2006-01-02")
}
func (this mytime) MustParseYMD(i string) time.Time {
	t, _ := this.ParseYMD(i)
	return t
}
func (this mytime) Parse(i string, format string) (time.Time, error) {
	return time.ParseInLocation(format, i, time.Local)
}
func (this mytime) MustParse(i string, format string) time.Time {
	t, _ := this.Parse(i, format)
	return t
}
