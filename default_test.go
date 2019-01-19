package logif

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

func TestDefaultLogger_Verbosity(t *testing.T) {
	type fields struct {
		verbosity int
		debug     int
		parent    *StdlibLogger
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"0",
			fields{0, 0, nil},
			0,
		},
		{
			"1",
			fields{1, 0, nil},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StdlibLogger{
				verbosity: tt.fields.verbosity,
				debug:     tt.fields.debug,
				parent:    tt.fields.parent,
			}
			DefaultLogger = s
			if got := Verbosity(); got != tt.want {
				t.Errorf("StdlibLogger.Verbosity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultLogger_SetVerbosity(t *testing.T) {
	type fields struct {
		verbosity int
		debug     int
		parent    *StdlibLogger
	}
	type args struct {
		v int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"0",
			fields{verbosity: 1, debug: 0, parent: nil},
			args{v: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StdlibLogger{
				verbosity: tt.fields.verbosity,
				debug:     tt.fields.debug,
				parent:    tt.fields.parent,
			}
			DefaultLogger = s
			SetVerbosity(tt.args.v)
			if DefaultLogger.verbosity != tt.args.v {
				t.Errorf("SetVerbosity did not set verbosity correctly")
			}
		})
	}
}

func TestDefaultLogger_Debugging(t *testing.T) {
	type fields struct {
		verbosity int
		debug     int
		parent    *StdlibLogger
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"0",
			fields{verbosity:0, debug:0, parent: nil},
			0,
		},
		{
			"1",
			fields{verbosity:0, debug:1, parent: nil},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StdlibLogger{
				verbosity: tt.fields.verbosity,
				debug:     tt.fields.debug,
				parent:    tt.fields.parent,
			}
			DefaultLogger = s
			if got := Debugging(); got != tt.want {
				t.Errorf("StdlibLogger.Debugging() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultLogger_SetDebugging(t *testing.T) {
	type fields struct {
		verbosity int
		debug     int
		parent    *StdlibLogger
	}
	type args struct {
		d int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"0",
			fields{verbosity: 1, debug: 0, parent: nil},
			args{d: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StdlibLogger{
				verbosity: tt.fields.verbosity,
				debug:     tt.fields.debug,
				parent:    tt.fields.parent,
			}
			DefaultLogger = s
			SetDebugging(tt.args.d)
			if DefaultLogger.verbosity != tt.args.d {
				t.Errorf("SetDebugging did not set debugging correctly")
			}
		})
	}
}

func TestDefaultLogger_Errorf(t *testing.T) {
	type fields struct {
		verbosity int
		debug     int
		parent    *StdlibLogger
		level int
	}
	type args struct {
		frmt string
		args []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		contain string
	}{
		{
			"no args",
			fields{0, 0, nil, LevelError},
			args{"no args", []interface{}{}},
			"no args",
		},
		{
			"2 args",
			fields{0, 0, nil, LevelError},
			args{"2 args %s %s", []interface{}{"a", "b"}},
			"2 args a b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StdlibLogger{
				verbosity: tt.fields.verbosity,
				debug:     tt.fields.debug,
				parent:    tt.fields.parent,
				level: tt.fields.level,
			}
			DefaultLogger = s
			var buf bytes.Buffer
			log.SetOutput(&buf)
			Errorf(tt.args.frmt, tt.args.args...)
			if !strings.Contains(buf.String(), tt.contain) {
				t.Errorf("did not get appropriate Errorf output: %s expected to contain %s", buf.String(), tt.contain)
			}
		})
	}
}

func TestDefaultLogger_Warningf(t *testing.T) {
	type fields struct {
		verbosity int
		debug     int
		parent    *StdlibLogger
		level int
	}
	type args struct {
		frmt string
		args []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		contain string
	}{
		{
			"no args",
			fields{0, 0, nil, LevelWarning},
			args{"no args", []interface{}{}},
			"no args",
		},
		{
			"2 args",
			fields{0, 0, nil, LevelWarning},
			args{"2 args %s %s", []interface{}{"a", "b"}},
			"2 args a b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StdlibLogger{
				verbosity: tt.fields.verbosity,
				debug:     tt.fields.debug,
				parent:    tt.fields.parent,
			}
			DefaultLogger = s
			var buf bytes.Buffer
			log.SetOutput(&buf)
			Warningf( tt.args.frmt, tt.args.args...)
			if !strings.Contains(buf.String(), tt.contain) {
				t.Errorf("did not get appropriate Warningf output: %s expected to contain %s", buf.String(), tt.contain)
			}
		})
	}
}

func TestDefaultLogger_Infof(t *testing.T) {
	type fields struct {
		verbosity int
		debug     int
		parent    *StdlibLogger
		level int
	}
	type args struct {
		frmt string
		args []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		contain string
	}{
		{
			"no args",
			fields{0, 0, nil, LevelInfo},
			args{"no args", []interface{}{}},
			"no args",
		},
		{
			"2 args",
			fields{0, 0, nil, LevelInfo},
			args{"2 args %s %s", []interface{}{"a", "b"}},
			"2 args a b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StdlibLogger{
				verbosity: tt.fields.verbosity,
				debug:     tt.fields.debug,
				parent:    tt.fields.parent,
			}
			DefaultLogger = s
			var buf bytes.Buffer
			log.SetOutput(&buf)
			Infof( tt.args.frmt, tt.args.args...)
			if !strings.Contains(buf.String(), tt.contain) {
				t.Errorf("did not get appropriate Infof output: %s expected to contain %s", buf.String(), tt.contain)
			}
		})
	}
}

func TestDefaultLogger_Debugf(t *testing.T) {
	type fields struct {
		verbosity int
		debug     int
		parent    *StdlibLogger
		level int
	}
	type args struct {
		frmt string
		args []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		contain string
	}{
		{
			"no args",
			fields{0, 0, nil, LevelDebug},
			args{"no args", []interface{}{}},
			"no args",
		},
		{
			"2 args",
			fields{0, 0, nil, LevelDebug},
			args{"2 args %s %s", []interface{}{"a", "b"}},
			"2 args a b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StdlibLogger{
				verbosity: tt.fields.verbosity,
				debug:     tt.fields.debug,
				parent:    tt.fields.parent,
			}
			DefaultLogger = s
			var buf bytes.Buffer
			log.SetOutput(&buf)
			Debugf( tt.args.frmt, tt.args.args...)
			if !strings.Contains(buf.String(), tt.contain) {
				t.Errorf("did not get appropriate Debugf output: %s expected to contain %s", buf.String(), tt.contain)
			}
		})
	}
}

func TestDefaultLogger_V(t *testing.T) {
	type fields struct {
		verbosity int
		debug     int
		parent    *StdlibLogger
		level int
	}
	type args struct {
		v int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		shouldOutput bool
	}{
		{
			"0 with 0 base",
			fields{0, 0, nil, LevelInfo},
			args{0},
			true,
		},
		{
			"1 with 0 base",
			fields{0, 0, nil, LevelInfo},
			args{1},
			false,
		},
		{
			"3 with 0 base",
			fields{0, 0, nil, LevelInfo},
			args{3},
			false,
		},
		{
			"3 with 3 base",
			fields{3, 0, nil, LevelInfo},
			args{3},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StdlibLogger{
				verbosity: tt.fields.verbosity,
				debug:     tt.fields.debug,
				parent:    tt.fields.parent,
				level: tt.fields.level,
			}
			DefaultLogger = s
			var buf bytes.Buffer
			log.SetOutput(&buf)
			V(tt.args.v).Infof("output")
			if tt.shouldOutput {
				if !strings.Contains(buf.String(), "output") {
					t.Errorf("StdlibLogger.V(%d).Infof(\"output\") did not print \"output\"", tt.args.v)
				}
			} else {
				if buf.String() != "" {
					t.Errorf("StdlibLogger.V(%d).Infof(...) should not have printed but did", tt.args.v)
				}
			}
		})
	}
}

func TestDefaultLogger_D(t *testing.T) {
	type fields struct {
		verbosity int
		debug     int
		parent    *StdlibLogger
		level int
	}
	type args struct {
		d int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		shouldOutput bool
	}{
		{
			"0 with 0 base",
			fields{0, 0, nil, LevelDebug},
			args{0},
			true,
		},
		{
			"1 with 0 base",
			fields{0, 0, nil, LevelDebug},
			args{1},
			false,
		},
		{
			"3 with 0 base",
			fields{0, 0, nil, LevelDebug},
			args{3},
			false,
		},
		{
			"3 with 3 base",
			fields{0, 3, nil, LevelDebug},
			args{3},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StdlibLogger{
				verbosity: tt.fields.verbosity,
				debug:     tt.fields.debug,
				parent:    tt.fields.parent,
				level: tt.fields.level,
			}
			DefaultLogger = s
			var buf bytes.Buffer
			log.SetOutput(&buf)
			D(tt.args.d).Debugf("output")
			if tt.shouldOutput {
				if !strings.Contains(buf.String(), "output") {
					t.Errorf("StdlibLogger.V(%d).Debugf(\"output\") did not print \"output\"", tt.args.d)
				}
			} else {
				if buf.String() != "" {
					t.Errorf("StdlibLogger.V(%d).Debugf(...) should not have printed but did", tt.args.d)
				}
			}
		})
	}
}

func TestDefaultLogger_SetLevel(t *testing.T) {
	s := &StdlibLogger{
		verbosity: 0,
		debug: 0,
		parent: nil,
		level: LevelDebug,
	}
	DefaultLogger = s
	tests := []struct{
		name string
		level int
		debug bool
		info bool
		warning bool
		error bool
	}{
		{"debug", LevelDebug, true, true, true, true},
		{"info", LevelInfo, false, true, true, true},
		{"warning", LevelWarning, false, false, true, true},
		{"error", LevelError, false, false, false, true},
		{"beyond error", 200, false, false, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetLevel(tt.level)
			var bufDebug bytes.Buffer
			log.SetOutput(&bufDebug)
			Debugf("output")
			if tt.debug {
				if !strings.Contains(bufDebug.String(), "output") {
					t.Errorf("Debugf output when level was too high")
				}
			}
			var bufInfo bytes.Buffer
			log.SetOutput(&bufInfo)
			Infof("output")
			if tt.info {
				if !strings.Contains(bufInfo.String(), "output") {
					t.Errorf("Infof output when level was too high")
				}
			}
			var bufWarning bytes.Buffer
			log.SetOutput(&bufWarning)
			Warningf("output")
			if tt.warning {
				if !strings.Contains(bufWarning.String(), "output") {
					t.Errorf("Warningf output when level was too high")
				}
			}
			var bufError bytes.Buffer
			log.SetOutput(&bufError)
			Errorf("output")
			if tt.error {
				if !strings.Contains(bufError.String(), "output") {
					t.Errorf("Errorf output when level was too high")
				}
			}
		})
	}
}