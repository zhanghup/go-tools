package tog

func Info(fmt string, f ...field) {
	toginfo.info(fmt, f...)
}

func Error(fmt string, f ...field) {
	toginfo.error(fmt, f...)
	togerr.error(fmt, f...)
}

func Warn(fmt string, f ...field) {
	toginfo.error(fmt, f...)
	togerr.error(fmt, f...)
}
