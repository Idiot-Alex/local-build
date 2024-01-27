package lblog

import (
	"fmt"
	"log"
)

// print error
func Error(err error) {
	if err == nil {
		return
	}

	log.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
}

// Info should be used to describe the example commands that are about to run.
func Info(format string, args ...interface{}) {
	log.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// Warning should be used to display a warning
func Warning(format string, args ...interface{}) {
	log.Printf("\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}
