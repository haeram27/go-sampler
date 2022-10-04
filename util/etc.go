package util

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

// avoid "not used" error
func AVOIDUNUSED(vals ...any) {
	for _, val := range vals {
		_ = val
	}
}

// avoid nil error
func AVOIDNIL(m any) any {
	if m == nil {
		return ""
	}
	return m
}

// if then else
func IfThenElse(condition bool, a any, b any) any {
	if condition {
		return a
	}
	return b
}

// check Windows OS type
func IsWindows() bool {
	return strings.EqualFold("windows", runtime.GOOS)
}

// path where application started at
func BasePath() string {
	path, err := os.Executable()
	if nil != err {
		return ""
	}

	return filepath.Dir(path)
}

// bin directory under BasePath()
func BinPath() string {
	return fmt.Sprintf("%s%cbin", BasePath(), os.PathSeparator)
}

// remove all empty spacesin string
func ConcatCharacters(data string) string {
	// \r  carriage return
	// \n  newline
	// \s  white space (U+0020)
	// \t  horizontal tab (U+0009)
	// \u3000  Unicode(double-byte) space (ideographic space, U+3000)
	regex := regexp.MustCompile(`\r?\n|\s+|\t+|\u3000+`)
	return regex.ReplaceAllString(data, "")
}
