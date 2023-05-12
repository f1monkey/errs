package errs

import (
	"fmt"
	"runtime"
	"strconv"
)

type Error struct {
	err   error
	stack []byte
}

func (e *Error) Error() string {
	return e.err.Error()
}

func (e *Error) Unwrap() error {
	return e.err
}

func (e *Error) StackTrace() []byte {
	return e.stack
}

func Errorf(format string, a ...any) *Error {
	var stack []byte
	for _, item := range a {
		if te, ok := item.(*Error); ok {
			stack = te.StackTrace()
		}
	}

	if stack == nil {
		stack = stacktrace(1, 10)
	}

	return &Error{
		err:   fmt.Errorf(format, a...),
		stack: stack,
	}
}

func stacktrace(skip int, num int) []byte {
	var buffer []byte

	pc := make([]uintptr, num)
	skip += 2 // skip runtime.Callers() and current function call
	n := runtime.Callers(skip, pc)
	cf := runtime.CallersFrames(pc[:n])
	for {
		frame, more := cf.Next()
		buffer = append(buffer, "\n"...)
		buffer = append(buffer, frame.Function...)
		buffer = append(buffer, "\n"...)
		buffer = append(buffer, "\t"...)
		buffer = append(buffer, frame.File...)
		buffer = append(buffer, ":"...)
		buffer = append(buffer, strconv.Itoa(frame.Line)...)

		if !more {
			break
		}
	}

	return buffer
}
