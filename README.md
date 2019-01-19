# `github.com/oko/logif`: Levelled Logging Interface

This Go package implements a levelled logging interface -- in other words, a logging interface that allows you to specify the "level" (importance) of a log message. This library is modelled after `glog` (one of the better-known levelled logging libraries in Go circles, made by Google), with two primary differences:

* **Separates interface and implementation:** the `logif.Logger` interface can be implemented by any chosen subsystem
* **No -vmodule or -tracelocation support:** these may be considered if I come across a need for them or a suitable PR is submitted

## License
This library is licensed under the [Apache 2.0 License](http://www.apache.org/licenses/LICENSE-2.0.txt).
