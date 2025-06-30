package common

import (
	"os"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)

	Root = filepath.Join(filepath.Dir(b)) // Insert the path to root in the second parameter.
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
