package pocketlog

// Level represents an available logging level
type Level byte

const (
	// LevelDebug represents the lowest level of log, mostly used
	LevelDebug Level = iota
	// LevelInfo represents a logging level that contains information
	LevelInfo
	// LevelError represents the highest logging level
	LevelError
)