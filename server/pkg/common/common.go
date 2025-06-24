package common

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetExecutableDir() (abspath string) {
	_, path, _, _ := runtime.Caller(1)
	path, err := filepath.Abs(path)
	if err == nil {
		abspath = filepath.Dir(path)
	}
	return abspath
}

func GetFirstArgument() string {
	if len(os.Args) < 2 {
		return ""
	}
	return os.Args[1]
}
