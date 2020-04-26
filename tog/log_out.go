package tog

func Info(fmt string, f ...Field) {
	togger.Info(fmt, f...)
}

func Error(fmt string, f ...Field) {
	togger.Error(fmt, f...)
}
