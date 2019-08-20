package logif

import (
	"errors"
	"strings"
)

// Logging level constants
const (
	LevelDebug = iota
	LevelInfo
	LevelWarning
	LevelError
)

// ErrNoSuchLogLevel is thrown when a parse
var ErrNoSuchLogLevel = errors.New("no such log level")

// LogLevelString returns the name representation of a given log level int
func LogLevelString(l int) (string, error) {
	switch l {
	case LevelDebug:
		return "DEBUG", nil
	case LevelInfo:
		return "INFO", nil
	case LevelWarning:
		return "WARN", nil
	case LevelError:
		return "ERROR", nil
	default:
		return "", ErrNoSuchLogLevel
	}
}

// ParseLogLevel parses a log level string into an int
func ParseLogLevel(l string) (int, error) {
	switch strings.ToUpper(l) {
	case "DEBUG":
		return LevelDebug, nil
	case "INFO":
		return LevelInfo, nil
	case "WARN":
		return LevelWarning, nil
	case "WARNING":
		return LevelWarning, nil
	case "ERR":
		return LevelError, nil
	case "ERROR":
		return LevelError, nil
	default:
		return -1, ErrNoSuchLogLevel
	}
}
