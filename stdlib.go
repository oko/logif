package logif

import (
	"fmt"
	"log"
)

type StdlibLogger struct {
	verbosity int
	debug int
	parent *StdlibLogger
	level int
}

func NewStdlibLogger() *StdlibLogger {
	return &StdlibLogger{
		verbosity: 0,
		debug: 0,
		parent: nil,
		level: LevelWarning,
	}
}

func (s *StdlibLogger) Verbosity() int {
	return s.verbosity
}

func (s *StdlibLogger) SetVerbosity(v int) {
	s.verbosity = v
}

func (s *StdlibLogger) Debugging() int {
	return s.debug
}

func (s *StdlibLogger) SetDebugging(d int) {
	s.debug = d
}

func (s *StdlibLogger) Errorf(frmt string, args ...interface{}) {
	// never actually suppress errors
	//if s.level > LevelError {
	//	return
	//}
	log.Printf(fmt.Sprintf("ERROR: %s", frmt), args...)
}

func (s *StdlibLogger) Warningf(frmt string, args ...interface{}) {
	if s.level > LevelWarning {
		return
	}
	log.Printf(fmt.Sprintf("WARN: %s", frmt), args...)
}

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

func (s *StdlibLogger) SetLevel(l int) {
	s.level = l
}

func (s *StdlibLogger) V(v int) Logger {
	return &StdlibLogger{
		verbosity: v,
		debug: s.debug,
		parent: s,
		level: s.level,
	}
}

func (s *StdlibLogger) D(d int) Logger {
	return &StdlibLogger{
		verbosity: s.verbosity,
		debug: d,
		parent: s,
		level: s.level,
	}
}