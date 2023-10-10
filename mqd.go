// Package mqd is used during development for printf-style debugging.  It
// shouldn't be in any release.
//
// All of the functions in this package log by default to stderr, regardless of
// how the log package's default logger is configured.
package mqd

/*
 * debug.go
 * Log debug messages
 * By J. Stuart McMurray
 * Created 20220829
 * Last Modified 20231010
 */

import (
	"fmt"
	"io"
	"log"
	"os"
)

// selog logs to stderr with microsecond precision.
var selog = log.New(os.Stderr, "", log.LstdFlags|log.Lmicroseconds)

// init notes the log package has been loaded.
func init() { log.Printf("MQD DEBUG PACKAGE LOADED") }

// SetOutput sets the output of the functions in this package to w.  By default
// the functions write to os.Stdout.
func SetOutput(w io.Writer) { selog.SetOutput(w) }

// Logf logs a message with the current file and line, if they can be
// determined.
func Logf(f string, a ...any) { nLogf(1, f, a...) }

// Fatalf is equivalent to Logf followed by os.Exit(code)
func Fatalf(code int, f string, a ...any) { nLogf(1, f, a...); os.Exit(code) }

// TODO prints a TODO: message with the current file and line, if they can be
// determined.
func TODO(f string, a ...any) { nLogf(1, "TODO: %s", fmt.Sprintf(f, a...)) }

// Here prints "Here" with the current file and line, if they can be
// determined, meant for wolf-fencing.
func Here() { nLogf(1, "Here") }

// Here1 prints "Here1" with the current file and line, if they can be
// determined, meant for wolf-fencing.
func Here1() { nLogf(1, "Here1") }

// Here2 prints "Here2" with the current file and line, if they can be
// determined, meant for wolf-fencing.
func Here2() { nLogf(1, "Here2") }

// DeleteThis prints "TODO: Delete this" with the current file and line, if
// they can be determined.
func DeleteThis() { nLogf(1, "TODO: Delete this") }

// Use prints "TODO: Use N variables" with the current file and line, if they
// can be determined.  N is the number of arguments.  This is intended to be
// used instead of _ = foo to silence warnings about temporarily-unused
// variables.
func Use(vars ...any) {
	if 1 == len(vars) {
		nLogf(1, "TODO: Use %d variable", len(vars))
	} else {
		nLogf(1, "TODO: Use %d variables", len(vars))
	}
}
