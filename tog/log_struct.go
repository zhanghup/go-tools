package tog

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

type Option struct {
	ShowLine   bool    `json:"show_line" yaml:"show_line"`
	LevelKey   string  `json:"level_key" yaml:"level_key"`
	TimeKey    string  `json:"time_key" yaml:"time_key"`
	LineEnding string  `json:"line_ending" yaml:"line_ending"`
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
	enable bool
	option *Option
	log    *zap.Logger
}

type Field struct {
	Name  string
	Value interface{}
}

func (this *Logger) Enable(flag bool) ILogger {
	this.enable = flag
	return this
}

func (this *Logger) fmt(f string, args ...interface{}) string {
	if len(args) == 0 {
		return f
	}
	return fmt.Sprintf(f, args...)
}

func (this *Logger) Info(f string, args ...interface{}) {
	if !this.enable {
		return
	}
	this.log.Info(this.fmt(f, args...))
}

func (this *Logger) Error(f string, args ...interface{}) {
	if !this.enable {
		return
	}
	this.log.Error(this.fmt(f, args...))
}

func (this *Logger) Warn(f string, args ...interface{}) {
	if !this.enable {
		return
	}
	this.log.Warn(this.fmt(f, args...))
}

func (this *Logger) Write(p []byte) (n int, err error) {
	if !this.enable {
		return
	}
	this.log.Info(string(p))
	return len(p), nil
}

func (this *Logger) setOption(option *Option) {
	this.option = option
	this.init()
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
		MessageKey:    "msg",
		LevelKey:      this.option.LevelKey,
		TimeKey:       this.option.TimeKey,
		NameKey:       "logger",
		CallerKey:     "line",
		StacktraceKey: "stacktrace",
		LineEnding:    this.option.LineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder, // 大写编码器
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder, // 一秒同步一次文件
		EncodeCaller:   this.option.EncodeCaller,
		EncodeName:     zapcore.FullNameEncoder,
	}

	encoder := zapcore.NewJSONEncoder(encoderConfig)
	if this.option.Type == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	core := zapcore.NewCore(
		encoder, // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		zap.NewAtomicLevelAt(zap.InfoLevel),                                             // 日志级别
	)

	options := make([]zap.Option, 0)
	if this.option.ShowLine {
		options = append(options, zap.AddCaller())
	}

	// 构造日志
	this.log = zap.New(
		core,
		options...,
	)
}

func NewLogger(opt *Option) ILogger {
	tog := &Logger{
		enable: true,
	}
	tog.setOption(opt)
	return tog
}
