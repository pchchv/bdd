package bdd

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

var curDir, _ = os.Getwd()

func getRelativeDirectory(targpath string) string {
	targpath = filepath.Clean(targpath)

	if fileName, err := filepath.Rel(curDir, targpath); err == nil && len(fileName) <= len(targpath) {
		targpath = fileName
	}

	return targpath
}

// MarkStackFull fills the output stack
func MarkStackFull() {
	for i := 1; ; i++ {
		s := SmarkStackFunc(i)
		if s == "" {
			break
		}
		fmt.Println(s)
	}
}

// MarkStack outputs the prefix stack line pos
func MarkStack(skip int, a ...interface{}) {
	fmt.Println(append([]interface{}{SmarkStack(skip + 1)}, a...)...)
}

// Mark outputs the prefix of the current line position
func Mark(a ...interface{}) {
	MarkStack(1, a...)
}

// Smark returns output prefix current line position
func Smark(a ...interface{}) string {
	return SmarkStack(1, a...)
}

// SmarkStack outputs stack information
func SmarkStack(skip int, a ...interface{}) string {
	_, fileName, line, ok := runtime.Caller(skip + 1)
	if !ok {
		return ""
	}

	fileName = getRelativeDirectory(fileName)

	return fmt.Sprintf("%s:%d %s", fileName, line, fmt.Sprint(a...))
}

// SmarkStackFunc outputs stack information
func SmarkStackFunc(skip int, a ...interface{}) string {
	pc, fileName, line, ok := runtime.Caller(skip + 1)
	if !ok {
		return ""
	}

	funcName := runtime.FuncForPC(pc).Name()
	funcName = filepath.Base(funcName)
	fileName = getRelativeDirectory(fileName)

	return fmt.Sprintf("%s:%d %s %s", fileName, line, funcName, fmt.Sprint(a...))
}
