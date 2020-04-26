package tog

import (
	"fmt"
	"go.uber.org/zap/zapcore"
	"runtime"
	"strings"
)

var toginfo *mylogger
var togerr *mylogger

func init() {
	toginfo = newLogger(&Option{
		Filename:   "./logs/stdout.log",
		MaxSize:    128,
		MaxBackups: 30,
		MaxAge:     7,
		Compress:   true,
	})

	togerr = newLogger(&Option{
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
	})
}
