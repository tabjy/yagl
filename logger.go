// Package YAGL, or, Yet Another Golang Logger, extends the golang built-in log
// package with level control, and multiple output io.Writer support. However,
// it's NOT compatible with built-in log packages, as log.Logger isn't exported
// as a interface, and many symbols and behaviours might differ.
// Package YAGL provides a Logger interface and a default implementation
// exported as yagl.Log. In addition, a set of static functions are provided to
// be called without creating a yagl.Log instance. (The underlying instance is
// created with default parameter.)
// Every log message is output on a separate line: if the message being
// printed does not end in a newline, the logger will add one.
// Package YAGL is designed to be thread-safe. A yagl.Logger instance can be
// used in multiple goroutines without precaution.
package yagl

// These flags define which text to prefix to each log entry generated.
// For example, flags 'Ldate | Ltime' (or LstdFlags) produce,
//		[INFO] 1970-01-01 00:00:00 LOG_MESSAGE
// while flags 'FlgDate | FlgTime | FlgMicroseconds | FlgLongFile | FlgPID' produce,
// 		[INFO] 1970-01-01 00:00:00.000000 /path/to/src/file.go:23 PID1234 LOG_MESSAGE
const (
	FlgDate         = 1 << iota         // the date in the local time zone: 2009/01/23
	FlgTime                             // the time in the local time zone: 01:23:23
	FlgMicroseconds                     // microsecond resolution: 01:23:23.123123. assumes Ltime.
	FlgLongFile                         // full file name and line number: /a/b/c/d.go:23
	FlgShortFile                        // final file name element and line number: d.go:23. overrides Llongfile
	FlgPID                              // process id of the goroutine calling logging function
	FlgUTC                              // if Ldate or Ltime is set, use UTC rather than the local time zone
	FlgStdFlags     = FlgDate | FlgTime // initial values for the standard logger
)

// These fields define logging level, each level is defined as a int and is a
// multiple of 10.
const (
	LvlTrace = 10 * iota
	LvlDebug
	LvlInfo
	LvlWarn
	LvlError
	LvlPanic
	LvlFatal
)

// Logger interface defines a minimum set of functions a well-designed logger
// instance should have. Any implementation of this interface MUST guarantees
// serialized access to a io.Writer, if called simultaneously from multiple
// goroutines.
// It's recommended to used the corresponding log level for a purpose described
// below.
type Logger interface {
	Trace(v ...interface{}) // lowest logging level, ideal for printing variable values during development.
	Debug(v ...interface{}) // anything that's too verbose to be included in "info" level.
	Info(v ...interface{})  // detailed regular logger operation, basically equivalent to built-in log.Logger.Print.
	Warn(v ...interface{})  // warning messages, which the application must be able to solve.
	Error(v ...interface{}) // error messages, this indicates the corresponding subroutine would return with a error soon.
	Panic(v ...interface{}) // something unexpected and unaccepted by design, the corresponding subroutine would throw a panic by calling this method.
	Fatal(v ...interface{}) // highest logger level, catastrophic failure, the entire process would exit with error code 1 by calling this method.

	Tracef(format string, v ...interface{}) // same as Trace(v ...interface{}), but with first argument being a format string.
	Debugf(format string, v ...interface{}) // same as Debug(v ...interface{}), but with first argument being a format string.
	Infof(format string, v ...interface{})  // same as Info(v ...interface{}), but with first argument being a format string.
	Warnf(format string, v ...interface{})  // same as Warn(v ...interface{}), but with first argument being a format string.
	Errorf(format string, v ...interface{}) // same as Error(v ...interface{}), but with first argument being a format string.
	Panicf(format string, v ...interface{}) // same as Panic(v ...interface{}), but with first argument being a format string.
	Fatalf(format string, v ...interface{}) // same as Fatal(v ...interface{}), but with first argument being a format string.
}