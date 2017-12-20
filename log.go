package yagl

import (
	"io"
	"sync"
	"bytes"
	"time"
	"fmt"
	"os"
	"runtime"
)

// Log is an implementation of the Logger interface for general purpose Logger.
type Log struct {
	mu    sync.Mutex  // ensures atomic writes; protects the following fields
	flag  int         // logging options
	level int         // logging levels
	out   []io.Writer // destinations for output
}

// New constructs a new Log instance which implements Logger. The flags
// parameter defines prefix of a line. The level parameter indicates to log all
// levels equal or equal to this level. The variadic parameter out sets the
// destination to which log data will be written.
// FlgLongFile and FlgShortFile flags can NOT be used together, or, a panic
// would be generated.
func New(flag int, level int, out ...io.Writer) Logger {
	// FlgLongFile and FlgShortFile can not be both set!
	if (flag&FlgLongFile != 0) && (flag&FlgShortFile != 0) {
		panic("yagl: FlgLongFile and FlgShortFile flag can not be both set")
	}

	return &Log{flag: flag, level: level, out: out}
}

// Level returns current logging level. Levels with value equal or higher to
// the returned number would be logged
func (l *Log) Level() int {
	return l.level
}

func (l *Log) buildHeader(builder *bytes.Buffer, level string, now time.Time) {
	builder.WriteString(fmt.Sprintf("[%s] ", level))

	if l.flag&(FlgDate|FlgTime) != 0 {
		if l.flag&FlgUTC != 0 {
			now = now.UTC()
		}
		if l.flag&FlgDate != 0 {
			year, month, day := now.Date()
			builder.WriteString(fmt.Sprintf("%04d-%02d-%02d ", year, month, day))
		}
		if l.flag&FlgTime != 0 {
			hour, min, sec := now.Clock()
			if l.flag&FlgMicroseconds != 0 {
				builder.WriteString(fmt.Sprintf("%02d:%02d:%02d.%02d ", hour, min, sec, now.Nanosecond()/1000))
			} else {
				builder.WriteString(fmt.Sprintf("%02d:%02d:%02d ", hour, min, sec))
			}
		}
		if l.flag&(FlgShortFile|FlgLongFile) != 0 {
			var ok bool
			_, file, line, ok := runtime.Caller(3) // Info() calls Output() calls buildHeader()
			if !ok {
				file = "???"
				line = 0
			}
			if l.flag&FlgShortFile != 0 {
				short := file
				for i := len(file) - 1; i > 0; i-- {
					if file[i] == '/' {
						short = file[i+1:]
						break
					}
				}
				file = short
			}

			builder.WriteString(fmt.Sprintf("%s:%d ", file, line))
		}
		if l.flag&FlgPID != 0 {
			builder.WriteString(fmt.Sprintf("PID%d ", os.Getpid()))
		}
	}
}

func (l *Log) output(level string, msg string) {
	now := time.Now()

	var builder bytes.Buffer
	l.buildHeader(&builder, level, now)

	builder.WriteString(msg)

	msgLen := len(msg)
	if msgLen == 0 || msg[msgLen-1] != '\n' {
		builder.WriteByte('\n')
	}

	res := builder.Bytes()

	l.mu.Lock()
	for _, dst := range l.out {
		dst.Write(res)
	}
	l.mu.Unlock()
}

// Trace is the lowest logging level. Arguments are handled in the manner of fmt.Print.
func (l *Log) Trace(v ...interface{}) {
	if LvlTrace >= l.level {
		l.output("TRACE", fmt.Sprint(v...))
	}
}

// Debug is for information too verbose to be included in "info" level. Arguments are handled in the manner of fmt.Print.
func (l *Log) Debug(v ...interface{}) {
	if LvlDebug >= l.level {
		l.output("DEBUG", fmt.Sprint(v...))
	}
}

// Info is the most commonly used logging level. Arguments are handled in the manner of fmt.Print.
func (l *Log) Info(v ...interface{}) {
	if LvlInfo >= l.level {
		l.output("INFO", fmt.Sprint(v...))
	}
}

// Warn indicates a warning. Arguments are handled in the manner of fmt.Print.
func (l *Log) Warn(v ...interface{}) {
	if LvlWarn >= l.level {
		l.output("WARN", fmt.Sprint(v...))
	}
}

// Error indicates an Error. Arguments are handled in the manner of fmt.Print.
func (l *Log) Error(v ...interface{}) {
	if LvlError >= l.level {
		l.output("ERROR", fmt.Sprint(v...))
	}
}

// Panic will call panic() after log is printed. Arguments are handled in the manner of fmt.Print.
func (l *Log) Panic(v ...interface{}) {
	msg := fmt.Sprint(v...)
	if LvlPanic >= l.level {
		l.output("PANIC", msg)
	}
	panic(msg)
}

// Fatal is the highest logging level with os.Exit(1) being called. Arguments are handled in the manner of fmt.Print.
func (l *Log) Fatal(v ...interface{}) {
	if LvlFatal >= l.level {
		l.output("FATAL", fmt.Sprint(v...))
	}
	os.Exit(1)
}

// Tracef is equivalent to Trace, but with arguments handled in the manner of fmt.Printf.
func (l *Log) Tracef(format string, v ...interface{}) {
	if LvlTrace >= l.level {
		l.output("TRACE", fmt.Sprintf(format, v...))
	}
}

// Debugf is equivalent to Debug, but with arguments handled in the manner of fmt.Printf.
func (l *Log) Debugf(format string, v ...interface{}) {
	if LvlDebug >= l.level {
		l.output("DEBUG", fmt.Sprintf(format, v...))
	}
}

// Infof is equivalent to Info, but with arguments handled in the manner of fmt.Printf.
func (l *Log) Infof(format string, v ...interface{}) {
	if LvlInfo >= l.level {
		l.output("INFO", fmt.Sprintf(format, v...))
	}
}

// Warnf is equivalent to Warn, but with arguments handled in the manner of fmt.Printf.
func (l *Log) Warnf(format string, v ...interface{}) {
	if LvlWarn >= l.level {
		l.output("WARN", fmt.Sprintf(format, v...))
	}
}

// Errorf is equivalent to Error, but with arguments handled in the manner of fmt.Printf.
func (l *Log) Errorf(format string, v ...interface{}) {
	if LvlError >= l.level {
		l.output("ERROR", fmt.Sprintf(format, v...))
	}
}

// Panicf is equivalent to Panic, but with arguments handled in the manner of fmt.Printf.
func (l *Log) Panicf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	if LvlPanic >= l.level {
		l.output("PANIC", msg)
	}
	panic(msg)
}

// Fatalf is equivalent to Fatal, but with arguments handled in the manner of fmt.Printf.
func (l *Log) Fatalf(format string, v ...interface{}) {
	if LvlFatal >= l.level {
		l.output("FATAL", fmt.Sprintf(format, v...))
	}
	os.Exit(1)
}
