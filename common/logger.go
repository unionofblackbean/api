package common

import (
	"log"
	"os"
)

const Lflags = log.Ldate | log.Ltime | log.Lshortfile | log.Lmsgprefix

func NewLogger(prefix string) *log.Logger {
	return log.New(os.Stdout, prefix, Lflags)
}
