package main

import (
	"os"

	"github.com/tabjy/yagl"
)

func main() {
	// First, we create a new logger instance.
	logger := yagl.New(
		// Flags that determines prefix before each line. None of them are mandatory, and we are almost using everything here
		yagl.FlgDate | yagl.FlgTime | yagl.FlgMicroseconds | yagl.FlgShortFile | yagl.FlgPID,
		// Desired logging level. Messages with logging level lower than this would not be logged
		yagl.LvlInfo, // LvlInfo looks good for production environment.
		os.Stdout, // Direct output content to standard out.
		// Optionally, we can add more output io.Writer after 3rd arg.
		/*
		func() io.Writer {
			fileWriter, _ := os.Create("/tmp/log.txt")
			return fileWriter
		}(),
		*/
		// And even more output io.Writer, thank to variadic parameter.
	)

	// Alternatively, we can use a standard logger in many cases
	// logger := yagl.StdLogger()
	// which is equivalent to yagl.New(yagl.FlgStdFlags, yagl.LvlInfo, os.Stdout)

	// Every log message is output on a separate line: if the message being printed does not end in a newline, the logger will add one.
	logger.Trace("Trace is the lowest logging level, ideal for printing variable values during development.")
	logger.Debug("Debug is for anything that's too verbose to be included in \"info\" level.")
	// However, these two line produce no output, as logging level is set to be yagl.LvlInfo

	logger.Info("Info is detailed regular logger operation, basically equivalent to built-in log.Logger.Print.")
	logger.Warn("Warn is for warning messages, which the application must be able to solve.")
	logger.Error("Error is for error messages, this indicates the corresponding subroutine would return with a error soon.")

	// This example will yield a panic soon, be ready to catch it.
	defer func() {
		r := recover()
		// Oh... Did I forget to mention that for every logging method, there is a alternated version with the first argument being a format string
		// So it behave like fmt.Printf
		// Infof for Info, Warnf for Warn, and Errorf for Error, etc...
		logger.Infof("Panic recovered: \n%v", r)

		// yagl.Fatal calls os.Exit(1)
		yagl.Fatal("Fatal is the highest logger level, catastrophic failure, the entire process would exit with error code 1 by calling this method.")
	}()

	// This cause a panic to be yielded.
	logger.Panic("Panic is for something unexpected and unaccepted by design, the corresponding subroutine would throw a panic by calling this method.")

	// A set of static methods are also provided:
	yagl.Trace("Like this")
	yagl.Debug("And this")
	yagl.Info("which is no different from calling yagl.StdLogger().Info")
	yagl.Warn("Just to make life for convenient")
	// ...

	// and ones with format string being the first argument
	yagl.Tracef("Answer to life, the universe, and everything: %d", 42)
	// Sadly, these lines produce nothing as a panic has already been yield...
}