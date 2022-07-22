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
func UNUSED(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}

// check nil when resolve data from empty interface
func GetNC(m interface{}) interface{} {
	if m == nil {
		return ""
	}
	return m
}

// if then else 문법 처리
func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

// Windows OS 여부 확인
func IsWindows() bool {
	return strings.EqualFold("windows", runtime.GOOS)
}

// 실행된 경로
func GetBasePath() string {
	path, err := os.Executable()
	if nil != err {
		return ""
	}

	return filepath.Dir(path)
}

// bin 폴더 경로
func GetBinPath() string {
	return fmt.Sprintf("%s%cbin", GetBasePath(), os.PathSeparator)
}

// 문자열에서 CRLF 및 모든 공백 제거
func OutputOmit(data string) string {
	regex := regexp.MustCompile(`\r?\n|\s+`)
	return regex.ReplaceAllString(data, "")
}
