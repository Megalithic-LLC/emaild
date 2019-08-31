package logger

import (
	"fmt"
	"log"
)

// Log level constants.
type Level int

const (
	LevelTrace Level = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

const (
	LevelTraceName = "TRACE"
	LevelDebugName = "DEBUG"
	LevelInfoName  = "INFO"
	LevelWarnName  = "WARN"
	LevelErrorName = "ERROR"
	LevelFatalName = "FATAL"
	LevelPanicName = "PANIC"
)

type Logger struct {
	isPanic bool
	isFatal bool
	isError bool
	isWarn  bool
	isInfo  bool
	isDebug bool
	isTrace bool
}

var logger *Logger

const noFormat = ""

func init() {
	logger = New()
}

func New() *Logger {
	return new(Logger)
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func (this *Logger) printf(debugLevelName string, format string, v ...interface{}) *Logger {
	var message string
	calldepth := 3
	if format == noFormat {
		v := append(v, 0)
		copy(v[1:], v[0:])
		v[0] = debugLevelName + " "
		message = fmt.Sprint(v...)
	} else {
		message = fmt.Sprintf(debugLevelName+" "+format, v...)
	}
	log.Output(calldepth, message)
	return this
}

// ----------------------------------------------------------------------------
// Public Setters
// ----------------------------------------------------------------------------

func SetLevel(level Level) *Logger { return logger.SetLevel(level) }
func (this *Logger) SetLevel(level Level) *Logger {
	this.isPanic = level <= LevelPanic
	this.isFatal = level <= LevelFatal
	this.isError = level <= LevelError
	this.isWarn = level <= LevelWarn
	this.isInfo = level <= LevelInfo
	this.isDebug = level <= LevelDebug
	this.isTrace = level <= LevelTrace
	return this
}

// ----------------------------------------------------------------------------
// Public IsXXX() routines
// ----------------------------------------------------------------------------

func IsPanic() bool { return logger.IsPanic() }
func (this *Logger) IsPanic() bool {
	return this.isPanic
}

func IsFatal() bool { return logger.IsFatal() }
func (this *Logger) IsFatal() bool {
	return this.isFatal
}

func IsError() bool { return logger.IsError() }
func (this *Logger) IsError() bool {
	return this.isError
}

func IsWarn() bool { return logger.IsWarn() }
func (this *Logger) IsWarn() bool {
	return this.isWarn
}

func IsInfo() bool { return logger.IsInfo() }
func (this *Logger) IsInfo() bool {
	return this.isInfo
}

func IsDebug() bool { return logger.IsDebug() }
func (this *Logger) IsDebug() bool {
	return this.isDebug
}

func IsTrace() bool { return logger.IsTrace() }
func (this *Logger) IsTrace() bool {
	return this.isTrace
}

// ----------------------------------------------------------------------------
// Public XXX() routines for leveled logging.
// ----------------------------------------------------------------------------

// --- Trace ------------------------------------------------------------------

func Trace(v ...interface{}) *Logger {
	if logger.IsTrace() {
		logger.printf(LevelTraceName, noFormat, v...)
	}
	return logger
}

func (this *Logger) Trace(v ...interface{}) *Logger {
	if this.isTrace {
		this.printf(LevelTraceName, noFormat, v...)
	}
	return this
}

func Tracef(format string, v ...interface{}) *Logger {
	if logger.IsTrace() {
		logger.printf(LevelTraceName, format, v...)
	}
	return logger
}

func (this *Logger) Tracef(format string, v ...interface{}) *Logger {
	if this.isTrace {
		this.printf(LevelTraceName, format, v...)
	}
	return this
}

// --- Debug ------------------------------------------------------------------

func Debug(v ...interface{}) *Logger {
	if logger.IsDebug() {
		logger.printf(LevelDebugName, noFormat, v...)
	}
	return logger
}

func (this *Logger) Debug(v ...interface{}) *Logger {
	if this.isDebug {
		this.printf(LevelDebugName, noFormat, v...)
	}
	return this
}

func Debugf(format string, v ...interface{}) *Logger {
	if logger.IsDebug() {
		logger.printf(LevelDebugName, format, v...)
	}
	return logger
}

func (this *Logger) Debugf(format string, v ...interface{}) *Logger {
	if this.isDebug {
		this.printf(LevelDebugName, format, v...)
	}
	return this
}

// --- Info -------------------------------------------------------------------

func Info(v ...interface{}) *Logger {
	if logger.IsInfo() {
		logger.printf(LevelInfoName, noFormat, v...)
	}
	return logger
}

func (this *Logger) Info(v ...interface{}) *Logger {
	if this.isInfo {
		this.printf(LevelInfoName, noFormat, v...)
	}
	return this
}

func Infof(format string, v ...interface{}) *Logger {
	if logger.IsInfo() {
		logger.printf(LevelInfoName, format, v...)
	}
	return logger
}

func (this *Logger) Infof(format string, v ...interface{}) *Logger {
	if this.isInfo {
		this.printf(LevelInfoName, format, v...)
	}
	return this
}

// --- Warn -------------------------------------------------------------------

func Warn(v ...interface{}) *Logger {
	if logger.IsWarn() {
		logger.printf(LevelWarnName, noFormat, v...)
	}
	return logger
}

func (this *Logger) Warn(v ...interface{}) *Logger {
	if this.isWarn {
		this.printf(LevelWarnName, noFormat, v...)
	}
	return this
}

func Warnf(format string, v ...interface{}) *Logger {
	if logger.IsWarn() {
		logger.printf(LevelWarnName, format, v...)
	}
	return logger
}

func (this *Logger) Warnf(format string, v ...interface{}) *Logger {
	if this.isWarn {
		this.printf(LevelWarnName, format, v...)
	}
	return this
}

// --- Error ------------------------------------------------------------------

func Error(v ...interface{}) *Logger {
	if logger.IsError() {
		logger.printf(LevelErrorName, noFormat, v...)
	}
	return logger
}

func (this *Logger) Error(v ...interface{}) *Logger {
	if this.isError {
		this.printf(LevelErrorName, noFormat, v...)
	}
	return this
}

func Errorf(format string, v ...interface{}) *Logger {
	if logger.IsError() {
		logger.printf(LevelErrorName, format, v...)
	}
	return logger
}

func (this *Logger) Errorf(format string, v ...interface{}) *Logger {
	if this.isError {
		this.printf(LevelErrorName, format, v...)
	}
	return this
}

// --- Fatal ------------------------------------------------------------------

func Fatal(v ...interface{}) *Logger {
	if logger.IsFatal() {
		logger.printf(LevelFatalName, noFormat, v...)
		log.Fatal("")
	}
	return logger
}

func (this *Logger) Fatal(v ...interface{}) *Logger {
	if this.isFatal {
		this.printf(LevelFatalName, noFormat, v...)
		log.Fatal("")
	}
	return this
}

func Fatalf(format string, v ...interface{}) *Logger {
	if logger.IsFatal() {
		logger.printf(LevelFatalName, format, v...)
		log.Fatal("")
	}
	return logger
}

func (this *Logger) Fatalf(format string, v ...interface{}) *Logger {
	if this.isFatal {
		this.printf(LevelFatalName, format, v...)
		log.Fatal("")
	}
	return this
}

// --- Panic ------------------------------------------------------------------

func Panic(v ...interface{}) *Logger {
	if logger.IsPanic() {
		logger.printf(LevelPanicName, noFormat, v...)
		log.Panic("")
	}
	return logger
}

func (this *Logger) Panic(v ...interface{}) *Logger {
	if this.isPanic {
		this.printf(LevelPanicName, noFormat, v...)
		log.Panic("")
	}
	return this
}

func Panicf(format string, v ...interface{}) *Logger {
	if logger.IsPanic() {
		logger.printf(LevelPanicName, format, v...)
		log.Panic("")
	}
	return logger
}

func (this *Logger) Panicf(format string, v ...interface{}) *Logger {
	if this.isPanic {
		this.printf(LevelPanicName, format, v...)
		log.Panic("")
	}
	return this
}
