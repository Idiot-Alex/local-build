package lblog

import (
	"fmt"
	"log"
)

// print error
func Error(format string, args ...interface{}) {
	log.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// Info should be used to describe the example commands that are about to run.
func Info(format string, args ...interface{}) {
	log.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// Warning should be used to display a warning
func Warning(format string, args ...interface{}) {
	log.Printf("\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}
