package tog

import (
	_ "embed"
	"github.com/zhanghup/go-tools"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

//go:embed config-default.yml
var defaultYamlConfig []byte

type Config struct {
	Enable bool `json:"enable" yaml:"enable"`

	LogPath      string `json:"log_path" yaml:"log_path"`
	LogLevel     Level  `json:"log_level" yaml:"log_level"`
	MaxSize      int    `json:"max_size" yaml:"max_size"`
	MaxAge       int    `json:"max_age" yaml:"max_age"`
	MaxBackups   int    `json:"max_backups" yaml:"max_backups"`
	Compress     bool   `json:"compress" yaml:"compress"`
	JsonFormat   bool   `json:"json_format" yaml:"json_format"`
	ShowLine     bool   `json:"show_line" yaml:"show_line"`
	LogInConsole bool   `json:"log_in_console" yaml:"log_in_console"`
	CallerSkip   int    `json:"caller_skip" yaml:"caller_skip"`
}

type Level string

const (
	LevelDebug Level = "debug"
	LevelInfo  Level = "info"
	LevelWarn  Level = "warn"
	LevelError Level = "error"
)

func NewLogger(configYaml ...[]byte) *Logger {
	cfg := struct {
		Log Config `json:"log" yaml:"log"`
	}{}
	err := tools.ConfOfByte(defaultYamlConfig, &cfg)
	if err != nil {
		panic(err)
	}

	for _, data := range configYaml {
		if data == nil {
			continue
		}

		err = tools.ConfOfByte(data, &cfg)
		if err != nil {
			panic(err)
		}
	}

	config := cfg.Log

	// 设置日志级别
	var level zapcore.Level
	switch config.LogLevel {
	case LevelDebug:
		level = zap.DebugLevel
	case LevelInfo:
		level = zap.InfoLevel
	case LevelWarn:
		level = zap.WarnLevel
	case LevelError:
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	// 定义日志切割配置
	hook := lumberjack.Logger{
		Filename:   config.LogPath,    // 日志文件的位置
		MaxSize:    config.MaxSize,    // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: config.MaxBackups, // 保留旧文件的最大个数
		Compress:   config.Compress,   // 是否压缩 disabled by default
		MaxAge:     config.MaxAge,     // days
	}

	// 判断是否控制台输出日志
	var syncer zapcore.WriteSyncer
	if config.LogInConsole {
		syncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook))
	} else {
		syncer = zapcore.AddSync(&hook)
	}

	// 定义zap配置信息
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "T",
		LevelKey:      "L",
		NameKey:       "default",
		CallerKey:     "C",
		MessageKey:    "M",
		StacktraceKey: "S",
		LineEnding:    zapcore.DefaultLineEnding,
		FunctionKey:   zapcore.OmitKey,
		EncodeTime: func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(time.Format("2006/01/02 15:04:05.000000"))
		}, // 自定义时间格式
		EncodeLevel: func(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString("[" + level.String() + "]")
		}, // 小写编码器
		EncodeCaller:   zapcore.ShortCallerEncoder, // 全路径编码器
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	var encoder zapcore.Encoder
	// 判断是否json格式输出
	if config.JsonFormat {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	core := zapcore.NewCore(
		encoder,
		syncer,
		level,
	)

	options := []zap.Option{
		zap.AddCallerSkip(config.CallerSkip),
		zap.AddStacktrace(zapcore.ErrorLevel),
	}

	// 判断是否显示代码行号
	if config.ShowLine {
		options = append(options, zap.AddCaller())
	}

	logger := zap.New(core, options...)

	return &Logger{
		enable: config.Enable,
		log:    logger,
		sugar:  logger.Sugar(),
		syncer: syncer,
	}
}
