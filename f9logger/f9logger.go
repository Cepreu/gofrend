package f9logger

import (
	"log"
	"os"
	"sync"
)

type f9Logger struct {
	*log.Logger
	filename string
}

var theLogger *f9Logger
var once sync.Once

//GetInstance creates a singleton instance of the GoFrend logger
func GetInstance() *f9Logger {
	once.Do(func() {
		theLogger = createLogger("f9Logger.log")
	})
	return theLogger
}

func createLogger(fname string) *f9Logger {
	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

	return &f9Logger{
		filename: fname,
		Logger:   log.New(file, "F9 ", log.Lshortfile),
	}
}
