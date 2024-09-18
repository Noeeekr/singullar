package projectpath

import (
	"path/filepath"
	"runtime"
)

// Useful to get the project root path.
// Path is the root path.

var (
	_, b, _, _ = runtime.Caller(0)

	Root = filepath.Join(filepath.Dir(b), "../..") // the second param is based on where this is located, should bring this to root
)
