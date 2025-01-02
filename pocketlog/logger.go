package pocketlog

import (
	"fmt"
	"io"
	"os"
)

// Logger is used to log information
type Logger struct {
	threshold Level
	output    io.Writer
	maxLength uint
}

// New returns a logger, ready to log at the required threshold.
// The default output is Stdout (Stderr for Errorf)
func New(threshold Level, opts ...Option) *Logger {
	logger := &Logger{threshold: threshold, output: os.Stdout, maxLength: 1000}
	for _, configFunc := range opts {
		configFunc(logger)
	}
	return logger
}

func (l *Logger) logf(lvl Level, format string, args ...any) {
	msg := fmt.Sprintf(format+"\n", args...)
	msgRune := []rune(msg)
	if uint(len(msgRune)) > l.maxLength {
		msg = fmt.Sprintf("%s...[TRIMMED]\n", string(msgRune[:l.maxLength]))
	} else {
		msg = string(msgRune)
	}
	_, _ = fmt.Fprintf(l.output, "%v %s", lvl, msg)
}

// Debug formats and prints a message if the log leve is debug
func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}
	if l.output == nil {
		l.output = os.Stdout
	}
	l.logf(LevelDebug, format, args...)
}

// Infof formats and prints a message if the leg level is info
func (l *Logger) Infof(format string, args ...any) {
	if l.threshold > LevelInfo {
		return
	}
	if l.output == nil {
		l.output = os.Stdout
	}
	l.logf(LevelInfo, format, args...)
}

// Errorf formats and prints a message if the log level is error
func (l *Logger) Errorf(format string, args ...any) {
	if l.threshold > LevelError {
		return
	}
	if l.output == nil {
		l.output = os.Stdout
	}
	l.logf(LevelError, format, args...)
}

func (l *Logger) Logf(lvl Level, format string, args ...any) {
	if l.threshold > lvl {
		return
	}
	l.logf(lvl, format, args)
}
