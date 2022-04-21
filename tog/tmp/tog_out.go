package tmp

import "github.com/zhanghup/go-tools"

func jsonFormat(o any, flag ...bool) string {
	s := tools.JSONString(o, flag...)
	if len(flag) > 0 && flag[0] {
		s = "\n" + s
	}
	return s
}

func Info(fmt string, args ...any) {
	Toginfo.Info(fmt, args...)
}

func InfoAsJson(o any, flag ...bool) {
	Toginfo.Info(jsonFormat(o, flag...))
}

func Error(fmt string, args ...any) {
	Toginfo.Error(fmt, args...)
	Togerr.Error(fmt, args...)
}

func ErrorAsJson(o any, flag ...bool) {
	Toginfo.Error(jsonFormat(o, flag...))
	Togerr.Error(jsonFormat(o, flag...))
}

func Warn(fmt string, args ...any) {
	Toginfo.Warn(fmt, args...)
}

func WarnAsJson(o any, flag ...bool) {
	Toginfo.Warn(jsonFormat(o, flag...))
}

func Enable(flag bool) {
	Toginfo.Enable(flag)
	Togerr.Enable(flag)
}
