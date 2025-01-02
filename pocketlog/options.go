package pocketlog

import (
	"io"
	"log"
)

// Option defines a functional option to our logger.
type Option func(*Logger)

// WithOutput returns a configuration function that sets the output.
func WithOutput(output io.Writer) Option {
	return func(l *Logger) {
		l.output = output
	}
}

// WithLength returns a configuration function that sets the maxLength of the resulting log,
// maxLength defaults to 1000 if not set with this function.
func WithLength(length uint) Option {
	if length < 10 {
		log.Fatalln("length should be at least 10")
	}
	return func(l *Logger) {
		l.maxLength = length
	}
}
