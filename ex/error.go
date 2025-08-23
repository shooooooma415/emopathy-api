package ex

import (
	"errors"
	"fmt"
	"runtime/debug"
	"strings"
)

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
		} else {
			return new(err, kind, append(converted.Args, args...)...)
		}
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

var ErrInRecover = errors.New("error in recover")

func Recover(e any) error {
	if e != nil {
		stack := debug.Stack()
		if err, ok := e.(error); ok {
			return Wrap(err, string(stack))
		} else {
			return Wrap(ErrInRecover, e, string(stack))
		}
	}
	return nil
}
