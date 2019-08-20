package logif

// Logger defines the interface for a logger which supports independently levelled info and debug logs
type Logger interface {
	// Errorf logs formatted errors
	Errorf(string, ...interface{})
	// Warningf logs formatted warnings
	Warningf(string, ...interface{})
	// Infof logs formatted info lines, and can be used for levelled logging in conjunction with the V(v int) method
	Infof(string, ...interface{})
	// Debugf logs formatted debug lines, and can be used for levelled logging in conjunction with the D(d int) method
	Debugf(string, ...interface{})
	// V returns a child logger that will only log INFO if its parent's level is less than or equal to its own
	V(int) Logger
	// D returns a child logger that will only log DEBUG if its parent's level is less than or equal to its own
	D(int) Logger
	// IsV returns whether verbosity is set at the given level
	IsV(int) bool
	// IsD returns whether debugging is set at the given level
	IsD(int) bool
	// Verbosity returns the current verbosity level for INFO messages
	Verbosity() int
	// Debugging returns the current debugging level for DEBUG messages
	Debugging() int
	// Level returns the current logging level
	Level() int
	// SetVerbosity sets the logger's verbosity level for INFO messages
	SetVerbosity(int)
	// SetVerbosity sets the logger's debugging level for DEBUG messages
	SetDebugging(int)
	// SetLevel sets the logger's base output level (DEBUG-INFO-WARN-ERROR)
	SetLevel(int)
}
