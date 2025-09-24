package util

import (
	"crypto/rand"
	"math/big"
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

func GenerateRandomStrings(size int) string {
	var s string = ""
	for i := len(s); i < size; {
		letter, err := rand.Int(rand.Reader, big.NewInt(91))
		if err != nil {
			panic(err)
		}
		if letter.Int64() > 64 {
			s += string(rune(letter.Int64()))
			i++
		}
	}
	return s
}
