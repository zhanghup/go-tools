package logger

import (
	"fmt"
	"go.uber.org/zap/zapcore"
	"runtime"
	"strings"
)

func OptionStdout() *Option {
	return &Option{
		Filename:   "./logs/stdout.log",
		MaxSize:    16,
		MaxBackups: 500,
		MaxAge:     60,
		LevelKey:   "level",
		TimeKey:    "time",
		ShowLine:   true,
		Compress:   true,
		Type:       "console",
		LineEnding: zapcore.DefaultLineEnding,
		EncodeCaller: func(c zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
			ss := strings.Split(c.FullPath(), "/")
			if len(ss) < 2 {
				return
			}
			enc.AppendString(strings.Join(ss[len(ss)-2:], "/"))
		},
	}
}

func OptionStderr() *Option {
	return &Option{
		Filename:   "./logs/stderr.log",
		MaxSize:    16,
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
