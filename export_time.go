package tools

import (
	"errors"
	"time"
)

type mytime struct {
	t   time.Time
	err error
}

var mytimeFormat = []string{
	"2006-01-02 15:04:05",
	"2006-01-02 15:04",
	"2006-01-02 15",
	"2006-01-02",

	"2006-1-2 15:04:05",
	"2006-1-2 15:04",
	"2006-1-2 15",
	"2006-1-2",

	"2006/01/02 15:04:05",
	"2006/01/02 15:04",
	"2006/01/02 15",
	"2006/01/02",

	"2006/1/2 15:04:05",
	"2006/1/2 15:04",
	"2006/1/2 15",
	"2006/1/2",

	"2006年01月02日 15时04分05秒",
	"2006年01月02日 15时04分",
	"2006年01月02日 15时",
	"2006年01月02日",
}

func TimeOfString(value string, format ...string) mytime {
	result := mytime{}
	if len(format) == 0 {
		format = mytimeFormat
	}

	for _, fmt := range format {
		v, err := time.ParseInLocation(fmt, value, time.Local)
		if err == nil {
			result.t = v
			return result
		}
	}

	result.t = time.Unix(0, 0)
	result.err = errors.New("[tools] 日期格式化失败")
	return result
}

func TimeOfUnix(value int64) mytime {
	result := mytime{}
	result.t = time.Unix(value, 0)
	return result
}

func Time(value ...time.Time) mytime {
	result := mytime{}

	if len(value) == 0 {
		result.t = time.Now()
	} else {
		result.t = value[0]
	}
	return result
}

func (this mytime) Time() time.Time {
	return this.t
}

func (this mytime) Unix() int64 {
	return this.t.Unix()
}

func (this mytime) String(format string) string {
	return this.t.Format(format)
}

func (this mytime) YMDHMS() string {
	return this.t.Format("2006-01-02 15:04:05")
}
func (this mytime) YMDHM() string {
	return this.t.Format("2006-01-02 15:04")
}
func (this mytime) YMDH() string {
	return this.t.Format("2006-01-02 15")
}
func (this mytime) YMD() string {
	return this.t.Format("2006-01-02")
}
func (this mytime) YM() string {
	return this.t.Format("2006-01")
}
func (this mytime) Y() string {
	return this.t.Format("2006")
}
func (this mytime) HMS() string {
	return this.t.Format("15:04:05")
}
func (this mytime) HM() string {
	return this.t.Format("15:04")
}
func (this mytime) H() string {
	return this.t.Format("15")
}
