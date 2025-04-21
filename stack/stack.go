package stack

import (
	"runtime"
	"strings"
)

// Frame wraps a runtime.Frame to provide some helper functions
type Frame struct {
	Function string
	File     string
	Line     int
}

// FuncName is the runtime.Frame.Function stripped down to just the function name
func (f *Frame) FuncName() string {
	name := f.Function
	i := strings.LastIndexByte(name, '.')
	return name[i+1:]
}

// FileName is the runtime.Frame.File stripped down to just the filename
func (f *Frame) FileName() string {
	name := f.File
	i := strings.LastIndexByte(name, '/')
	return name[i+1:]
}

// Stack returns a stack Frame
func Stack() *Frame {
	return SkipLevel(1)
}

// SkipLevel returns a stack Frame skipping the number of supplied frames.
// This is primarily used by other libraries who use this package
// internally as the additional.
func SkipLevel(skip int) *Frame {
	var frame [3]uintptr
	runtime.Callers(skip+2, frame[:])
	frames := runtime.CallersFrames(frame[:])
	v, _ := frames.Next()

	f := &Frame{
		Function: v.Function,
		File:     v.File,
		Line:     v.Line,
	}
	return f
}
