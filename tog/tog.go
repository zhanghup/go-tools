package tog

import (
	"fmt"
	"github.com/zhanghup/go-tools/tog/logger"
	"go.uber.org/zap/zapcore"
	"runtime"
	"strings"
)

var toginfo logger.ILogger
var togerr logger.ILogger

func init() {
	info := logger.OptionStdout()
	info.EncodeCaller = func(c zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		strs := []string{}
		for i := 7; i < 8; i++ {
			_, str, l, ok := runtime.Caller(i)
			if ok {
				ss := strings.Split(str,"/")
				if len(ss) > 3{
					str = strings.Join(ss[len(ss) - 3:],"/")
				}
				strs = append(strs, fmt.Sprintf("%s:%d", str, l))
			} else {
				break
			}
		}
		enc.AppendString(strings.Join(strs, "") + "")
	}
	toginfo = logger.NewLogger(info)

	err := logger.OptionStderr()
	err.EncodeCaller = func(c zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
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
	}
	togerr = logger.NewLogger(err)

}
