package logif

import (
	"log"
	"os"
	"testing"
)

func TestExample(t *testing.T) {
	log.SetOutput(os.Stdout)
	Debugf("hello %s", "debug")
	Infof("hello %s", "info")
	Warningf("hello %s", "warning")
	Errorf("hello %s", "error")

	DefaultLogger.Debugf("hello %s", "dl debug")
	DefaultLogger.Infof("hello %s", "dl info")
	DefaultLogger.Warningf("hello %s", "dl warning")
	DefaultLogger.Errorf("hello %s", "dl error")

	DefaultLogger.SetLevel(LevelDebug)
	DefaultLogger.SetVerbosity(1)
	DefaultLogger.V(1).Infof("level 1")
	DefaultLogger.V(2).Infof("level 2 not shown")

	if DefaultLogger.IsV(2) {
		log.Printf("this line shouldn't print")
		DefaultLogger.V(2).Infof("level 2 not shown")
	}

	DefaultLogger.SetVerbosity(3)
	if DefaultLogger.IsV(2) {
		log.Printf("this line should print")
		DefaultLogger.V(2).Infof("level 2 shown when verbosity set to 3")
	}
}
