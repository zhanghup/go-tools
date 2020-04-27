package logger

type ILogger interface {
	Info(message string, f ...Field)
	Error(message string, f ...Field)
	Warn(message string, f ...Field)
}
