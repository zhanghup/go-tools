package tog

import "github.com/zhanghup/go-tools"

func jsonFormat(o interface{}, flag ...bool) string{
	s := tools.Str.JSONString(o, flag...)
	if len(flag) > 0 && flag[0]{
		s = "\n" +s
	}
	return s
}

func Info(fmt string, args ...interface{}) {
	toginfo.Info(fmt, args...)
}

func InfoAsJson(o interface{}, flag ...bool) {
	toginfo.Info(jsonFormat(o,flag...))
}

func Error(fmt string, args ...interface{}) {
	toginfo.Error(fmt, args...)
	togerr.Error(fmt, args...)
}

func ErrorAsJson(o interface{}, flag ...bool) {
	toginfo.Error(jsonFormat(o,flag...))
	togerr.Error(jsonFormat(o,flag...))
}

func Warn(fmt string, args ...interface{}) {
	toginfo.Warn(fmt, args...)
}

func WarnAsJson(o interface{}, flag ...bool) {
	toginfo.Warn(jsonFormat(o,flag...))
}
