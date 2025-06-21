package log

import (
	"log"
	"os"
	"sync"
)

var (
	LogErr  *log.Logger
	LogInfo *log.Logger
)

var once sync.Once

func InitLoggers() {

	once.Do(func() {
		LogErr = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)
		LogInfo = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	})

}
