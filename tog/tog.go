package tog

import (
	"fmt"
	"go.uber.org/zap/zapcore"
	"runtime"
	"strings"
)

var Toginfo ILogger
var Togerr ILogger

func init() {
	Toginfo = NewLogger(&Option{
		Filename:   "./logs/stdout.log",
		MaxSize:    120,
		MaxBackups: 500,
		MaxAge:     60,
		LevelKey:   "level",
		TimeKey:    "time",
		ShowLine:   true,
		Compress:   true,
		Type:       "console",
		LineEnding: zapcore.DefaultLineEnding,
		EncodeCaller: func(c zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
			strs := []string{}
			for i := 7; i < 8; i++ {
				_, str, l, ok := runtime.Caller(i)
				if ok {
					ss := strings.Split(str,"/")
					if len(ss) > 1{
						str = strings.Join(ss[len(ss) - 1:],"/")
					}
					strs = append(strs, fmt.Sprintf("%s:%d", str, l))
				} else {
					break
				}
			}
			enc.AppendString(strings.Join(strs, "") + "")
		},
	})

	Togerr = NewLogger(&Option{
		Filename:   "./logs/stderr.log",
		MaxSize:    20,
		MaxBackups: 500,
		MaxAge:     60,
		LevelKey:   "level",
		TimeKey:    "time",
		ShowLine:   true,
		Compress:   true,
		Type:       "console",
		LineEnding: zapcore.DefaultLineEnding,
		EncodeCaller: func(c zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
			strs := []string{}
			for i := 0; i < 100; i++ {
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
