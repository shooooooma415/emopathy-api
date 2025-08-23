package ex

import (
	"errors"
	"fmt"
	"runtime/debug"
	"strings"
)

// Error represents an extended error with additional context
type Error struct {
	Err   error     `json:"error,omitempty"`
	Stack []string  `json:"stack,omitempty"`
	Args  []any     `json:"args,omitempty"`
	Kind  ErrorKind `json:"kind,omitempty"`
}

var _ error = (*Error)(nil)

func (e *Error) Error() string {
	if len(e.Args) > 0 {
		return fmt.Sprintf("%s %+v", e.Err.Error(), e.Args)
	}
	return e.Err.Error()
}

func new(err error, kind ErrorKind, args ...any) *Error {
	return &Error{
		Err:   err,
		Stack: stackTrace(),
		Args:  args,
		Kind:  kind,
	}
}

// Wrap wraps an error with additional arguments
func Wrap(err error, args ...any) error {
	return wrap(err, emptyKind, args...)
}

func wrap(err error, kind ErrorKind, args ...any) error {
	if err == nil {
		return nil
	}

	if isMultiple(err) {
		return new(err, kind, args...)
	}

	converted := &Error{}
	if ok := errors.As(err, &converted); ok {
		if converted.Kind == kind {
			converted.Args = append(converted.Args, args...)
			return converted
		}
		return new(err, kind, append(converted.Args, args...)...)
	}

	return new(err, kind, args...)
}

func (e *Error) Unwrap() error {
	return e.Err
}

func isMultiple(err error) bool {
	var multi interface{ Unwrap() []error }
	return errors.As(err, &multi)
}

func stackTrace() []string {
	stack := debug.Stack()
	lines := strings.Split(string(stack), "\n")
	if len(lines) > 2 {
		return lines[2:]
	}
	return lines
}

// ErrInRecover is returned when a panic is recovered but the recovered value is not an error
var ErrInRecover = errors.New("error in recover")

// Recover recovers from a panic and converts it to an error
func Recover(e any) error {
	if e != nil {
		stack := debug.Stack()
		if err, ok := e.(error); ok {
			return Wrap(err, string(stack))
		}
		return Wrap(ErrInRecover, e, string(stack))
	}
	return nil
}
