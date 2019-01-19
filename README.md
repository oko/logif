# `github.com/oko/logif`: Levelled Logging Interface

[![Build Status](https://travis-ci.org/oko/logif.svg?branch=master)](https://travis-ci.org/oko/logif)
[![Go Report Card](https://goreportcard.com/badge/github.com/oko/logif)](https://goreportcard.com/report/github.com/oko/logif)
[![codecov](https://codecov.io/gh/oko/logif/branch/master/graph/badge.svg)](https://codecov.io/gh/oko/logif)

This Go package implements a levelled logging interface -- in other words, a logging interface that allows you to specify the "level" (importance) of a log message. This library is modelled after `glog` (one of the better-known levelled logging libraries in Go circles, made by Google), with two primary differences:

* **Separates interface and implementation:** the `logif.Logger` interface can be implemented by any chosen subsystem
* **No -vmodule or -tracelocation support:** these may be considered if I come across a need for them or a suitable PR is submitted

## Quick Start

The library ships with a default logger preconfigured based on the standard library's `log` package:

```
package main

import (
	"github.com/oko/logif"
)

func main() {
	logif.Errorf("Hello world!")
	logif.SetLevel(logif.LevelDebug)
	logif.Debugf("%#v", logif.DefaultLogger)
}
```

You can change the default logger (which is used by `logif`'s top-level Debugf/Infof/Warningf/Errorf functions) by setting `logif.DefaultLogger` to a type implementing the `logif.Logger` interface.

## License
This library is licensed under the [Apache 2.0 License](http://www.apache.org/licenses/LICENSE-2.0.txt).
