package logif

import "testing"

type logLevelCase struct {
	input int
	out   string
	err   error
}
type parseLevelCase struct {
	input string
	out   int
	err   error
}

func TestLogLevelString(t *testing.T) {
	input := []logLevelCase{
		{LevelDebug, "DEBUG", nil},
		{LevelInfo, "INFO", nil},
		{LevelWarning, "WARN", nil},
		{LevelError, "ERROR", nil},
		{-1, "", ErrNoSuchLogLevel},
	}
	for _, i := range input {
		out, err := LogLevelString(i.input)
		if out != i.out || err != i.err {
			t.Errorf("LogLevelString(%d) returned %s, %s, expected %s, %s", i.input, out, err, i.out, i.err)
		}
	}
}

func TestParseLogLevel(t *testing.T) {
	input := []parseLevelCase{
		{"DEBUG", LevelDebug, nil},
		{"INFO", LevelInfo, nil},
		{"WARN", LevelWarning, nil},
		{"WARNING", LevelWarning, nil},
		{"ERR", LevelError, nil},
		{"ERROR", LevelError, nil},
		{"MISSING", -1, ErrNoSuchLogLevel},
	}
	for _, i := range input {
		out, err := ParseLogLevel(i.input)
		if out != i.out || err != i.err {
			t.Errorf("ParseLogLevel(%s) returned %d, %s, expected %d, %s", i.input, out, err, i.out, i.err)
		}
	}
}
