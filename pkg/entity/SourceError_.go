package entity

import (
	"fmt"
	"log"
	"runtime"
)

// SourceError will get additional error information, useful for creating a stackable
// error.
type SourceError struct {
	// Func is a function signature along with it's argument, created when calling
	//  new(entity.StackError).With(err, args...)
	Func string

	// File is a filename of the caller, created when calling
	//  new(entity.StackError).With(err, args...)
	File string

	// Line is the line number of the caller, created when calling
	//  new(entity.StackError).With(err, args...)
	Line int

	// skip is the number of stack frames to skip, created when calling
	//  new(entity.StackError).With(err, args...)
	skip int

	// err is the underlying error
	err error
}

// Unwrap get underlying error.
func (e *SourceError) Unwrap() error { return e.err }

// Error implement error interface.
func (e *SourceError) Error() string {
	if e == e.err {
		e.err = ErrRecursive
	}

	return fmt.Sprintf("\n%s\n%s:%d\n\t%s", e.Func, e.File, e.Line, e.err)
}

func (e *SourceError) Skip(skip int) *SourceError {
	e.skip = skip

	return e
}

// With create a new.
func (e *SourceError) With(err error, args ...interface{}) *SourceError {
	if err == nil {
		return nil
	}

	if skip := 1; e.skip < skip {
		e.skip = skip
	}

	pc, file, line, _ := runtime.Caller(e.skip)
	argsV := ""

	for _, arg := range args {
		if argsV != "" {
			argsV += ", "
		}
		argsV += fmt.Sprintf("%T(%#v)", arg, arg)
	}

	fn := fmt.Sprintf("%s(%s)", runtime.FuncForPC(pc).Name(), argsV)
	*e = SourceError{fn, file, line, e.skip, err}

	log.Printf("Error: %v | Func: %s | File: %s | Line: %d", err, e.Func, e.File, e.Line)

	return e
}
