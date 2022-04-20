package tog

type ILogger interface {
	Info(fmt string, args ...any)
	Error(fmt string, args ...any)
	Warn(fmt string, args ...any)

	Write(p []byte) (n int, err error)
	Enable(flag bool) ILogger
}
