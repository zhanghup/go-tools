package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
	"time"
)

type Option struct {
	Filename   string  `json:"filename" yaml:"filename"`
	MaxSize    int     `json:"maxsize" yaml:"maxsize"`
	MaxAge     int     `json:"maxage" yaml:"maxage"`
	MaxBackups int     `json:"maxbackups" yaml:"maxbackups"`
	LocalTime  bool    `json:"localtime" yaml:"localtime"`
	Compress   bool    `json:"compress" yaml:"compress"`
	Type       string  `json:"type" yaml:"type"` // json or console
	Field      []Field `json:"field" yaml:"field"`

	EncodeCaller func(c zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) `json:"-" yaml:"-"`
}

type Logger struct {
	option *Option
	log    *zap.Logger
}

type Field struct {
	Name  string
	Value interface{}
}

func (this *Logger) Info(message string, f ...Field) {
	this.log.Info(message, getField(f...)...)
}

func (this *Logger) Error(message string, f ...Field) {
	this.log.Error(message, getField(f...)...)
}

func (this *Logger) Warn(message string, f ...Field) {
	this.log.Warn(message, getField(f...)...)
}

func (this *Logger) setOption(option *Option) {
	if option.EncodeCaller == nil {
		option.EncodeCaller = func(c zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
			ss := strings.Split(c.FullPath(), "/")
			if len(ss) < 2 {
				return
			}
			enc.AppendString(strings.Join(ss[len(ss)-2:], "/"))
		}
	}

	this.option = option
	this.init()
}

func (this *Logger) Fields(f ...Field) {
	this.log = this.log.With(getField(f...)...)
}

func (this *Logger) init() {
	hook := lumberjack.Logger{
		Filename:   this.option.Filename,   // 日志文件路径
		MaxSize:    this.option.MaxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: this.option.MaxBackups, // 日志文件最多保存多少个备份
		MaxAge:     this.option.MaxAge,     // 文件最多保存多少天
		Compress:   this.option.Compress,   // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "line",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder, // 大写编码器
		EncodeTime:     func(t time.Time, enc zapcore.PrimitiveArrayEncoder) { enc.AppendString(t.Format("2006-01-02 15:04:05.000")) },
		EncodeDuration: zapcore.SecondsDurationEncoder, // 一秒同步一次文件
		EncodeCaller:   this.option.EncodeCaller,
		EncodeName:     zapcore.FullNameEncoder,
	}

	encoder := zapcore.NewJSONEncoder(encoderConfig)
	if this.option.Type == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	core := zapcore.NewCore(
		encoder,                                                                         // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		zap.NewAtomicLevelAt(zap.InfoLevel),                                             // 日志级别
	)

	// 构造日志
	this.log = zap.New(
		core,
		zap.AddCaller(),
		zap.Development(),
		zap.Fields(getField(this.option.Field...)...),
	)
	this.log.With()

}

func getField(f ...Field) []zapcore.Field {
	fs := make([]zap.Field, 0)
	for _, o := range f {
		fs = append(fs, zap.Any(o.Name, o.Value))
	}
	return fs
}

func NewLogger(opt *Option) ILogger {
	tog := &Logger{}
	tog.setOption(opt)
	return tog
}
