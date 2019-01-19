package logif

import "testing"

func TestExample(t *testing.T) {
	Debugf("hello %s", "user")
	Infof("hello %s", "user")
	Warningf("hello %s", "user")
	Errorf("hello %s", "user")

	DefaultLogger.Debugf("hello %s", "user")
	DefaultLogger.Infof("hello %s", "user")
	DefaultLogger.Warningf("hello %s", "user")
	DefaultLogger.Errorf("hello %s", "user")
}
