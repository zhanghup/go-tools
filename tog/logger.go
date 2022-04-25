package tog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	enable bool
	log    *zap.Logger
	sugar  *zap.SugaredLogger
	syncer zapcore.WriteSyncer
}

func (this *Logger) Debug(args ...any) {
	if !this.enable {
		return
	}
	this.sugar.Debug(args...)
}
func (this *Logger) Info(args ...any) {
	if !this.enable {
		return
	}
	this.sugar.Info(args...)
}
func (this *Logger) Warn(args ...any) {
	if !this.enable {
		return
	}
	this.sugar.Warn(args...)
}
func (this *Logger) Error(args ...any) {
	if !this.enable {
		return
	}
	this.sugar.Error(args...)
}

func (this *Logger) Debugf(template string, args ...any) {
	if !this.enable {
		return
	}
	this.sugar.Debugf(template, args...)
}
func (this *Logger) Infof(template string, args ...any) {
	if !this.enable {
		return
	}
	this.sugar.Infof(template, args...)
}
func (this *Logger) Warnf(template string, args ...any) {
	if !this.enable {
		return
	}
	this.sugar.Warnf(template, args...)
}
func (this *Logger) Errorf(template string, args ...any) {
	if !this.enable {
		return
	}
	this.sugar.Errorf(template, args...)
}

func (this *Logger) Debugs(msg string, fields ...zap.Field) {
	if !this.enable {
		return
	}
	this.log.Debug(msg, fields...)
}
func (this *Logger) Infos(msg string, fields ...zap.Field) {
	if !this.enable {
		return
	}
	this.log.Info(msg, fields...)
}
func (this *Logger) Warns(msg string, fields ...zap.Field) {
	if !this.enable {
		return
	}
	this.log.Warn(msg, fields...)
}
func (this *Logger) Errors(msg string, fields ...zap.Field) {
	if !this.enable {
		return
	}
	this.log.Error(msg, fields...)
}

func (this *Logger) Write(data []byte) (int, error) {
	if !this.enable {
		return 0, nil
	}
	return this.syncer.Write(data)
}
