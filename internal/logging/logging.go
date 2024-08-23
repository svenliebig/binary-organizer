package logging

import (
	"fmt"
)

type Level int

const (
	TraceLevel Level = iota
	DebugLevel
	InfoLevel
	ErrorLevel
	NoneLevel
)

func (l Level) String() string {
	return []string{"TRACE", "DEBUG", "INFO", "ERROR"}[l]
}

var level Level = NoneLevel

func SetLevel(l Level) {
	level = l
}

// Trace logs a message at the trace severity level.
func Trace(args ...any) {
	log(TraceLevel, args...)
}

// Debug logs a message at the debug severity level.
func Debug(args ...any) {
	log(DebugLevel, args...)
}

// Info logs a message at the info severity level.
func Info(args ...any) {
	log(InfoLevel, args...)
}

// Error logs a message at the error severity level.
func Error(args ...any) {
	log(ErrorLevel, args...)
}

// Fn logs the entry and exit of a function with the given name.
// the logging is done at the trace severity level.
//
// Example:
//
//	func foo() {
//		logging.Fn("foo")
//		// ...
//	}
//
// Output:
//
//	[TRACE] entering foo
//	[TRACE] exiting foo
func Fn(name string) {
	Trace("entering", name)
	defer Trace("exiting", name)
}

func log(l Level, args ...any) {
	if l < level {
		return
	}

	prefix := fmt.Sprintf("[%s]", l.String())
	fmt.Println(prefix, fmt.Sprint(args...))
}
