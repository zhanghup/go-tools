package tog

import (
	"go.uber.org/zap"
	"io"
)

var _logger *Logger
var Writer io.Writer

func Init(configYaml ...[]byte) {
	_logger = NewLogger(configYaml...)
	Writer = _logger
}

func init() {
	Init(nil)
}

func Debug(args ...any) {
	_logger.Debug(args...)
}
func Info(args ...any) {
	_logger.Info(args...)
}
func Warn(args ...any) {
	_logger.Warn(args...)
}
func Error(args ...any) {
	_logger.Error(args...)
}

func Debugf(template string, args ...any) {
	_logger.Debugf(template, args...)
}
func Infof(template string, args ...any) {
	_logger.Infof(template, args...)
}
func Warnf(template string, args ...any) {
	_logger.Warnf(template, args...)
}
func Errorf(template string, args ...any) {
	_logger.Errorf(template, args...)
}

func Debugs(msg string, fields ...zap.Field) {
	_logger.Debugs(msg, fields...)
}
func Infos(msg string, fields ...zap.Field) {
	_logger.Infos(msg, fields...)
}
func Warns(msg string, fields ...zap.Field) {
	_logger.Warns(msg, fields...)
}
func Errors(msg string, fields ...zap.Field) {
	_logger.Errors(msg, fields...)
}

func Write(data []byte) (int, error) {
	return _logger.Write(data)
}
