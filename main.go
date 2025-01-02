package main

import (
	"os"

	"github.com/daniel-ojo-williams/logger/pocketlog"
)

func main() {
	logger := pocketlog.New(pocketlog.LevelDebug, pocketlog.WithOutput(os.Stdout), pocketlog.WithLength(30))

	logger.Infof("A little copying is  better than a little dependency")
	logger.Debugf("Make the zero (%d) value useful", 0)
}
