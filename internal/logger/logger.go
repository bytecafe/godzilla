// Package logger is a logging package with log level support.
package logger

import (
	"io"
	"log" // nolint:depguard
	"path"
	"runtime"
)

// LogLevel for the default logger
type LogLevel int

const (
	LError LogLevel = iota
	LWarning
	LInfo
	LDebug
)

// SetLevel set the minmal LogLevel for the default logger
func SetLevel(l LogLevel) {
	defaultLogger.level = l
}

// SetOutput redirets log message into `w` instead of os.Stderr
func SetOutput(w io.Writer) {
	defaultLogger.SetOutput(w)
}

// Error logs error messages.
func Error(f string, args ...interface{}) {
	logf(LError, f, args...)
}

// Warn logs warning messages.
func Warn(f string, args ...interface{}) {
	logf(LWarning, f, args...)
}

// Info logs some readable info.
func Info(f string, args ...interface{}) {
	logf(LInfo, f, args...)
}

// Debug logs some verbose messsages for debug purpose.
func Debug(f string, args ...interface{}) {
	logf(LDebug, f, args...)
}

func logf(level LogLevel, f string, args ...interface{}) {
	// if the target level is greater than the configured level, just ignore it
	if level > defaultLogger.level {
		return
	}

	prefix := ""
	switch level {
	case LError:
		prefix += "[ERROR]"
	case LWarning:
		prefix += "[WARN]"
	case LInfo:
		prefix += "[INFO]"
	case LDebug:
		prefix += "[DEBUG]"
	}

	_, file, _, _ := runtime.Caller(2) // nolint
	pkg := path.Base(path.Dir(file))
	prefix += "[" + pkg + "] "

	defaultLogger.Printf(prefix+f, args...)
}

type loggerWraper struct {
	*log.Logger

	level LogLevel
}

var defaultLogger = loggerWraper{
	Logger: log.Default(),
	level:  LInfo,
}
