package tog_test

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"testing"
	"time"
)

func TestZap(t *testing.T) {
	cfg := zap.NewProductionConfig()
	cfg.Encoding = "console"

	(func() {
		_, err := os.Open("logs")
		if os.IsNotExist(err) {
			_ = os.Mkdir("logs", os.ModePerm)
		}
	})()

	config := cfg.EncoderConfig
	config.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	config.EncodeLevel = func(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString("[" + strings.ToUpper(level.String()) + "]")
	}
	cfg.EncoderConfig = config
	cfg.OutputPaths = []string{"./logs/stdout.log"}
	cfg.ErrorOutputPaths = []string{"./logs/stderr.log"}

	fmt.Println(tools.JSONString(cfg.EncoderConfig))

	logger, _ := cfg.Build()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	go func() {
		for {
			sugar.Infof("SUCCESS to fetch URL: %s", "url")
			sugar.Errorf("FAILED to fetch URL: %s", "url")
		}
	}()
	time.Sleep(time.Second * 60)
}
