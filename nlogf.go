package mqd

/*
 * nlogf.go
 * Core logf function underlying nearly everything else wraps
 * By J. Stuart McMurray
 * Created 20230415
 * Last Modified 20230415
 */

import (
	"fmt"
	"runtime"
)

// nLogf logs a message, but allows setting the depth of the stack frame from
// which to get the file and line.
func nLogf(n int, f string, a ...any) {
	selog.Printf(
		"[%s] %s",
		fileAndLineTag(n+1),
		fmt.Sprintf(f, a...),
	)
}

// fileAndLineTag returns a tag with the current file and line, if they can
// be determined, otherwise "Unknown location".  n will be passed to
// runtime.Caller and should be the number of functions between fileAndLineTag
// and the calling function of interest.
func fileAndLineTag(n int) string {
	_, fn, ln, ok := runtime.Caller(n + 1)
	if !ok {
		return "Unknown location"
	}
	return fmt.Sprintf("%s:%d", fn, ln)
}
