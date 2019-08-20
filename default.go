package logif

// DefaultLogger stores the Logger implementation used by the logif top level functions
var DefaultLogger Logger = NewStdlibLogger()

// Errorf logs a formatted message at ERROR level
func Errorf(fmt string, args ...interface{}) {
	DefaultLogger.Errorf(fmt, args...)
}

// Warningf logs a formatted message at WARN level
func Warningf(fmt string, args ...interface{}) {
	DefaultLogger.Warningf(fmt, args...)
}

// Infof logs a formatted message at INFO level
func Infof(fmt string, args ...interface{}) {
	DefaultLogger.Infof(fmt, args...)
}

// Debugf logs a formatted message at DEBUG level
func Debugf(fmt string, args ...interface{}) {
	DefaultLogger.Debugf(fmt, args...)
}

// V returns a sublogger that only logs INFO messages if the parent verbosity is greater than v
func V(v int) Logger {
	return DefaultLogger.V(v)
}

// D returns a sublogger that only logs DEBUG messages if the parent debugging verbosity is greater than d
func D(d int) Logger {
	return DefaultLogger.D(d)
}

// IsV returns whether verbosity is set at the given level
func IsV(v int) bool {
	return DefaultLogger.IsV(v)
}

// IsD returns whether debugging is set at the given level
func IsD(d int) bool {
	return DefaultLogger.IsD(d)
}

// Verbosity returns the default logger verbosity
func Verbosity() int {
	return DefaultLogger.Verbosity()
}

// SetVerbosity sets the default logger verbosity
func SetVerbosity(v int) {
	DefaultLogger.SetVerbosity(v)
}

// Debugging returns the default logger debugging verbosity
func Debugging() int {
	return DefaultLogger.Debugging()
}

// SetDebugging sets the default logger debugging verbosity
func SetDebugging(d int) {
	DefaultLogger.SetDebugging(d)
}

// Level returns the default logger's logging level
func Level() int {
	return DefaultLogger.Level()
}

// SetLevel sets the default logger's overall level threshold
func SetLevel(l int) {
	DefaultLogger.SetLevel(l)
}
