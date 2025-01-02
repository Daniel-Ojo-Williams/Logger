package pocketlog_test

import (
	"fmt"
	"testing"

	"github.com/daniel-ojo-williams/logger/pocketlog"
)

const (
	debugMessage = "Why write I still all one, ever the same"
	inforMessage = "And keep invention in a noted weed"
	errorMessage = "That every word doth almost tell my name"
)

func TestLogger_DebugInfoErrof(t *testing.T) {
	type testCase struct {
		level    pocketlog.Level
		expected string
	}

	tt := map[string]testCase{
		"debug": {
			level:    pocketlog.LevelDebug,
			expected: "0 " + debugMessage + "\n" + "1 " + inforMessage + "\n" + "2 " + errorMessage + "\n",
		},
		"info": {
			level:    pocketlog.LevelInfo,
			expected: "1 " + inforMessage + "\n" + "2 " + errorMessage + "\n",
		},
		"error": {
			level:    pocketlog.LevelError,
			expected: "2 " + errorMessage + "\n",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}
			tesetedLogger := pocketlog.New(tc.level, pocketlog.WithOutput(tw))
			tesetedLogger.Debugf(debugMessage)
			tesetedLogger.Infof(inforMessage)
			tesetedLogger.Errorf(errorMessage)

			if tc.expected != tw.contents {
				t.Fatalf("invalid contents, expected %q, got %q", tc.expected, tw.contents)
			}
		})
	}
}

func TestLogger_OutputLength(t *testing.T) {
	maxLength := uint(11)
	expected := "0 Why write I...[TRIMMED]\n"
	tw := &testWriter{}
	logger := pocketlog.New(pocketlog.LevelDebug, pocketlog.WithOutput(tw), pocketlog.WithLength(maxLength))
	logger.Debugf(debugMessage)

	if tw.contents != expected {
		t.Fatalf("invalid length, expected %s, got %s", expected, tw.contents)
	}
}

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug)
	debugLogger.Debugf("Hello, %s", "world")
	// Output: 0 Hello, world
}

// testWriter is a struct that implements io.Writer
// We use it to validate that we can write to a specific output.
type testWriter struct {
	contents string
}

func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.contents = tw.contents + string(p)
	return len(p), nil
}
