package logif

var DefaultLogger = NewStdlibLogger()

func Errorf(fmt string, args ...interface{}) {
	DefaultLogger.Errorf(fmt, args...)
}

func Warningf(fmt string, args ...interface{}) {
	DefaultLogger.Warningf(fmt, args...)
}

func Infof(fmt string, args ...interface{}) {
	DefaultLogger.Infof(fmt, args...)
}

func Debugf(fmt string, args ...interface{}) {
	DefaultLogger.Debugf(fmt, args...)
}

func V(v int) Logger {
	return DefaultLogger.V(v)
}

func D(d int) Logger {
	return DefaultLogger.D(d)
}

func Verbosity() int {
	return DefaultLogger.Verbosity()
}

func SetVerbosity(v int) {
	DefaultLogger.SetVerbosity(v)
}

func Debugging() int {
	return DefaultLogger.Debugging()
}

func SetDebugging(d int) {
	DefaultLogger.SetDebugging(d)
}

func SetLevel(l int) {
	DefaultLogger.SetLevel(l)
}