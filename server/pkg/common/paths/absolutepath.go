package paths

import (
	"path/filepath"
	"runtime"
)

// Useful to get the project root path.
// Path is the root path.

var (
	_, b, _, _ = runtime.Caller(0)

	Root = filepath.Join(filepath.Dir(b), "../") // Insert the path to root in the second parameter.
)
