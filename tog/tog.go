package tog

import (
	"fmt"
	"github.com/zhanghup/go-tools/tog/logger"
	"go.uber.org/zap/zapcore"
	"runtime"
	"strings"
)

var toginfo = NewStdoutLog(nil)
var togerr = NewStderrLog(nil)

func NewStdoutLog(opt *logger.Option) logger.ILogger {
	var o *logger.Option
	if opt != nil {
		o = opt
	} else {
		o = &logger.Option{
			Filename:   "./logs/stdout.log",
			MaxSize:    128,
			MaxBackups: 30,
			MaxAge:     7,
			Compress:   true,
		}
	}
	return logger.NewLogger(o)
}

func NewStderrLog(opt *logger.Option) logger.ILogger {
	var o *logger.Option
	if opt != nil {
		o = opt
	} else {
		o = &logger.Option{
			Filename:   "./logs/stderr.log",
			MaxSize:    128,
			MaxBackups: 30,
			MaxAge:     7,
			Compress:   true,
			Type:       "console",
			EncodeCaller: func(c zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
				strs := []string{}
				for i := 7; i < 100; i++ {
					_, str, l, ok := runtime.Caller(i)
					if ok {
						strs = append(strs, fmt.Sprintf("\n\t%s:%d", str, l))
					} else {
						break
					}
				}
				enc.AppendString(strings.Join(strs, "") + "\n")
			},
		}
	}
	return logger.NewLogger(o)
}

func WidthField(extra ...map[string]interface{}) {
	toginfo.Fields(getField(extra...)...)
	togerr.Fields(getField(extra...)...)
}
