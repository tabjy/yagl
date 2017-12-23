package yagl

import "os"

var std = New(FlgStdFlags, LvlInfo, os.Stderr)

// StdLogger return a standard logger equivalent to
// 		New(FlgStdFlags, LvlInfo, os.Stdout)
func StdLogger() Logger {
	return std
}

// Trace is equivalent to calling std.Trace(), with std being a standard logger.
func Trace(v ...interface{}) { std.Trace(v...) }

// Debug is equivalent to calling std.Debug(), with std being a standard logger.
func Debug(v ...interface{}) { std.Debug(v...) }

// Info is equivalent to calling std.Info(), with std being a standard logger.
func Info(v ...interface{}) { std.Info(v...) }

// Warn is equivalent to calling std.Warn(), with std being a standard logger.
func Warn(v ...interface{}) { std.Warn(v...) }

// Error is equivalent to calling std.Error(), with std being a standard logger.
func Error(v ...interface{}) { std.Error(v...) }

// Panic is equivalent to calling std.Panic(), with std being a standard logger.
func Panic(v ...interface{}) { std.Panic(v...) }

// Fatal is equivalent to calling std.Fatal(), with std being a standard logger.
func Fatal(v ...interface{}) { std.Fatal(v...) }

// Tracef is equivalent to calling std.Tracef(), with std being a standard logger.
func Tracef(format string, v ...interface{}) { std.Tracef(format, v...) }

// Debugf is equivalent to calling std.Debugf(), with std being a standard logger.
func Debugf(format string, v ...interface{}) { std.Debugf(format, v...) }

// Infof is equivalent to calling std.Infof(), with std being a standard logger.
func Infof(format string, v ...interface{}) { std.Infof(format, v...) }

// Warnf is equivalent to calling std.Warnf(), with std being a standard logger.
func Warnf(format string, v ...interface{}) { std.Warnf(format, v...) }

// Errorf is equivalent to calling std.Errorf(), with std being a standard logger.
func Errorf(format string, v ...interface{}) { std.Errorf(format, v...) }

// Panicf is equivalent to calling std.Panicf(), with std being a standard logger.
func Panicf(format string, v ...interface{}) { std.Panicf(format, v...) }

// Fatalf is equivalent to calling std.Fatalf(), with std being a standard logger.
func Fatalf(format string, v ...interface{}) { std.Fatalf(format, v...) }
