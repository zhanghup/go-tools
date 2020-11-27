package tog

type ILogger interface {
	Info(fmt string, args ...interface{})
	Error(fmt string, args ...interface{})
	Warn(fmt string, args ...interface{})

	Write(p []byte) (n int, err error)
}
