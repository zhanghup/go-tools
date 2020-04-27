package tog

import "github.com/zhanghup/go-tools/tog/logger"

func Info(fmt string, extra ...map[string]interface{}) {
	toginfo.Info(fmt, getField(extra...)...)
}

func Error(fmt string, extra ...map[string]interface{}) {
	toginfo.Error(fmt, getField(extra...)...)
	togerr.Error(fmt, getField(extra...)...)
}

func Warn(fmt string, extra ...map[string]interface{}) {
	toginfo.Warn(fmt, getField(extra...)...)
}

func getField(extra ...map[string]interface{}) []logger.Field {
	result := make([]logger.Field, 0)
	for _, o := range extra {
		for k, v := range o {
			result = append(result, logger.Field{Name: k, Value: v})
		}
	}
	return result
}
