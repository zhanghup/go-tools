package tog

import "go.uber.org/zap"

type Logger struct {
	option *zap.Config
	log    *zap.Logger
}

func NewLogger(config *zap.Config) *Logger {
	return &Logger{option: config}
}

var infoLog *Logger
var errorLog *Logger
