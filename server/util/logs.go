package util

import (
	"log"
	"os"
)

var Error = log.New(os.Stdout, "[ERROR] ", log.Llongfile+log.Ldate+log.Ltime)
var Info = log.New(os.Stdout, "[INFO] ", log.Ldate+log.Ltime)
