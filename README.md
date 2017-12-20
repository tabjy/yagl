# YAGL: Yet Another Golang Logger

Package YAGL, or, Yet Another Golang Logger, extends the golang built-in log package with level control, and multiple output io.Writer support.

## Getting Started

Please note that YAGL is NOT a drop-in replacement for the Golang built-in log package as log.Logger isn't exported as a interface (but a struct), and many symbols and behaviours might differ.

### Installing

The Go tool chain has made everything so simple. To install the latest version, just do:
```bash
$ go get github.com/tabjy/yagl
```

### Importing

Importing YAGL is no different from importing other packages:
```go
import "github.com/tabjy/yagl"
```
Then the packages is imported as `yagl`.

## Basic Usage

```go
logger := yagl.StdLogger()
logger.Info("Great! YAGL is working perfectly!")
```

This would output something like
```bash
[INFO] 2017-12-20 02:57:18 Great! YAGL is working perfectly!
```

A [well-explained example](example/basics.go) is also available.

## Documentation

The codes come together with inline documentation.

### Logging Levels
6 logging level are provided:
- **Trace**:  lowest logging level, ideal for printing variable values during development.
- **Debug**: anything that's too verbose to be included in "info" level.
- **Info**: detailed regular logger operation, basically equivalent to built-in log.Logger.Print.
- **Warn**: warning messages, which the application must be able to solve.
- **Error**: error messages, this indicates the corresponding subroutine would return with a error soon.
- **Panic**: something unexpected and unaccepted by design, the corresponding subroutine would throw a panic by calling this method.
- **Fatal**: highest logger level, catastrophic failure, the entire process would exit with error code 1 by calling this method.

## TODOs

- [ ] color code different logging levels.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.