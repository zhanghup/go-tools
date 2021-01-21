package tog

import "github.com/zhanghup/go-tools"

func jsonFormat(o interface{}, flag ...bool) string{
	s := tools.JSONString(o, flag...)
	if len(flag) > 0 && flag[0]{
		s = "\n" +s
	}
	return s
}

func Info(fmt string, args ...interface{}) {
	Toginfo.Info(fmt, args...)
}

func InfoAsJson(o interface{}, flag ...bool) {
	Toginfo.Info(jsonFormat(o,flag...))
}

func Error(fmt string, args ...interface{}) {
	Toginfo.Error(fmt, args...)
	Togerr.Error(fmt, args...)
}

func ErrorAsJson(o interface{}, flag ...bool) {
	Toginfo.Error(jsonFormat(o,flag...))
	Togerr.Error(jsonFormat(o,flag...))
}

func Warn(fmt string, args ...interface{}) {
	Toginfo.Warn(fmt, args...)
}

func WarnAsJson(o interface{}, flag ...bool) {
	Toginfo.Warn(jsonFormat(o,flag...))
}
