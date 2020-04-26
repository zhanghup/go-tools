package tog

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"runtime"
	"strings"
	"time"
)

type Option struct {
	Filename   string `json:"filename" yaml:"filename"`
	MaxSize    int    `json:"maxsize" yaml:"maxsize"`
	MaxAge     int    `json:"maxage" yaml:"maxage"`
	MaxBackups int    `json:"maxbackups" yaml:"maxbackups"`
	LocalTime  bool   `json:"localtime" yaml:"localtime"`
	Compress   bool   `json:"compress" yaml:"compress"`

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

func (this *Logger) Info(fmt string, f ...Field) {
	fs := make([]zap.Field, 0)
	for _, o := range f {
		fs = append(fs, zap.Any(o.Name, o.Value))
	}
	this.log.Info(fmt, fs...)
}

func (this *Logger) Error(fmt string, f ...Field) {
	fs := make([]zap.Field, 0)
	for _, o := range f {
		fs = append(fs, zap.Any(o.Name, o.Value))
	}
	this.log.Error(fmt, fs...)
}

func (this *Logger) SetOption(option *Option) {
	if option.EncodeCaller == nil {
		option.EncodeCaller = func(c zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
			strs := make([]string, 0)
			for i := 6; i < 7; i++ {
				_, str, l, ok := runtime.Caller(i)
				if ok {
					strs = append(strs, fmt.Sprintf("%s:%d", str, l))
				} else {
					break
				}
			}
			enc.AppendString(strings.Join(strs, ","))
		}
	}

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

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		zap.NewAtomicLevelAt(zap.InfoLevel),                                             // 日志级别
	)

	// 构造日志
	this.log = zap.New(
		core,
		zap.AddCaller(),
		zap.Development(),
		zap.Fields(
			//zap.String("serviceName", "serviceName"),
		),
	)

}

var togger *Logger

func init() {
	togger = &Logger{}
	togger.SetOption(&Option{
		Filename:   "./logs/info.log",
		MaxSize:    128,
		MaxBackups: 30,
		MaxAge:     7,
		Compress:   true,
	})
}
