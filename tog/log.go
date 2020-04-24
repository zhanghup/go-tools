package tog

import "go.uber.org/zap"

func initConfig() zap.Config {
	return zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       true,
		DisableCaller:     false,
		DisableStacktrace: false,
	}
}

func Init(config zap.Config) {

}

func init() {
	cfg := initConfig()
	cfg.Build()
}
