package logif

import (
	"fmt"
	"log"
)

// StdlibLogger is a Logger implementation that utilizes the Go standard library log package
type StdlibLogger struct {
	verbosity int
	debug     int
	parent    *StdlibLogger
	level     int
}

// NewStdlibLogger returns a new logger with level defaulted to WARN
func NewStdlibLogger() *StdlibLogger {
	return &StdlibLogger{
		verbosity: 0,
		debug:     0,
		parent:    nil,
		level:     LevelWarning,
	}
}

// Verbosity returns the default logger verbosity
func (s *StdlibLogger) Verbosity() int {
	return s.verbosity
}

// SetVerbosity sets the default logger verbosity
func (s *StdlibLogger) SetVerbosity(v int) {
	s.verbosity = v
}

// Debugging returns the default logger debugging verbosity
func (s *StdlibLogger) Debugging() int {
	return s.debug
}

// SetDebugging sets the default logger debugging verbosity
func (s *StdlibLogger) SetDebugging(d int) {
	s.debug = d
}

// Errorf logs a formatted message at ERROR level
func (s *StdlibLogger) Errorf(frmt string, args ...interface{}) {
	// never actually suppress errors
	//if s.level > LevelError {
	//	return
	//}
	log.Printf(fmt.Sprintf("ERROR: %s", frmt), args...)
}

// Warningf logs a formatted message at WARN level
func (s *StdlibLogger) Warningf(frmt string, args ...interface{}) {
	if s.level > LevelWarning {
		return
	}
	log.Printf(fmt.Sprintf("WARN: %s", frmt), args...)
}

// Infof logs a formatted message at INFO level
func (s *StdlibLogger) Infof(frmt string, args ...interface{}) {
	if s.level > LevelInfo {
		return
	}
	if s.parent != nil {
		if s.verbosity <= s.parent.verbosity {
			log.Printf(fmt.Sprintf("INFO[%d]: %s", s.verbosity, frmt), args...)
		}
	} else {
		log.Printf(fmt.Sprintf("INFO: %s", frmt), args...)
	}
}

// Debugf logs a formatted message at DEBUG level
func (s *StdlibLogger) Debugf(frmt string, args ...interface{}) {
	if s.level > LevelDebug {
		return
	}
	if s.parent != nil {
		if s.debug <= s.parent.debug {
			log.Printf(fmt.Sprintf("DEBUG[%d]: %s", s.debug, frmt), args...)
		}
	} else {
		log.Printf(fmt.Sprintf("DEBUG: %s", frmt), args...)
	}
}

// SetLevel sets the default logger's overall level threshold
func (s *StdlibLogger) SetLevel(l int) {
	s.level = l
}

// V returns a sublogger that only logs INFO messages if the parent verbosity is greater than v
func (s *StdlibLogger) V(v int) Logger {
	return &StdlibLogger{
		verbosity: v,
		debug:     s.debug,
		parent:    s,
		level:     s.level,
	}
}

// D returns a sublogger that only logs DEBUG messages if the parent debugging verbosity is greater than d
func (s *StdlibLogger) D(d int) Logger {
	return &StdlibLogger{
		verbosity: s.verbosity,
		debug:     d,
		parent:    s,
		level:     s.level,
	}
}
