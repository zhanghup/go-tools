package tog_test

import (
	_ "embed"
	"github.com/zhanghup/go-tools/tog"
	"go.uber.org/zap"
	"os"
	"testing"
)

func TestZap(t *testing.T) {
	for i := 0; i < 10; i++ {
		tog.Infof("SUCCESS to fetch URL: %s", "url")
		tog.Errorf("FAILED to fetch URL: %s", "url")
	}

	//go func() {
	//	for {
	//		tog.Infof("SUCCESS to fetch URL: %s", "url")
	//		tog.Errorf("FAILED to fetch URL: %s", "url")
	//	}
	//}()
	//
	//time.Sleep(time.Second * 10)
}

func TestZap2(t *testing.T) {
	cfg := zap.NewProductionConfig()
	cfg.Encoding = "console"

	(func() {
		_, err := os.Open("logs")
		if os.IsNotExist(err) {
			_ = os.Mkdir("logs", os.ModePerm)
		}
	})()

	logger, _ := cfg.Build()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	//go func() {
	//	for {
	//		sugar.Infof("SUCCESS to fetch URL: %s", "url")
	//		sugar.Errorf("FAILED to fetch URL: %s", "url")
	//	}
	//}()
	//time.Sleep(time.Second * 60)

	for i := 0; i < 10; i++ {
		sugar.Infof("SUCCESS to fetch URL: %s", "url")
		sugar.Errorf("FAILED to fetch URL: %s", "url")
	}

}
