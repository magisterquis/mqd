package mqd

/*
 * debug.go
 * Log debug messages
 * By J. Stuart McMurray
 * Created 20230513
 * Last Modified 20230513
 */

import (
	"fmt"
	"runtime"
)

// finishThisMsg usually indicates something ought to be finished.
const finishThisMsg = "TODO: Finish This"

// FinishThisError notes the file and line where ErrFinishThis was called.
type FinishThisError struct {
	File string
	Line int
}

// Error implements the error interface.  It returns "TODO: Finish this" with
// err.File and err.Line, if set.
func (err FinishThisError) Error() string {
	if "" == err.File || 0 == err.Line {
		return finishThisMsg
	}
	return fmt.Sprintf("%s (%s:%d)", finishThisMsg, err.File, err.Line)
}

// ErrFinishThis returns a FinishThisError with the current file and line,
// if available.
func ErrFinishThis() FinishThisError {
	_, f, l, ok := runtime.Caller(1)
	if !ok {
		return FinishThisError{}
	}
	return FinishThisError{
		File: f,
		Line: l,
	}
}

// FinishThis prints "TODO: Finish this" with the current file and line, if
// they can be determined.
func FinishThis() { nLogf(1, "TODO: Finish this") }
