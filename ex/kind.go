package ex

import (
	"errors"
)

type ErrorKind string

const (
	emptyKind               ErrorKind = ""
	InvalidArgumentError    ErrorKind = "InvalidArgument"
	UnauthenticatedError    ErrorKind = "Unauthenticated"
	AccessDeniedError       ErrorKind = "AccessDenied"
	NotFoundError           ErrorKind = "NotFound"
	TimeoutError            ErrorKind = "Timeout"
	ConflictError           ErrorKind = "Conflict"
	ResourceExhaustedError  ErrorKind = "ResourceExhausted"
	CanceledError           ErrorKind = "Canceled"
	InternalError           ErrorKind = "Internal"
	DependentServiceError   ErrorKind = "DependentService"
	FailedPreconditionError ErrorKind = "FailedPrecondition"
)

func (e *Error) Is(target error) bool {
	if target == nil {
		return false
	}

	if e == target { //nolint: err113 // errors.Is() を使えという指摘だが、ここでは e そのものを意図的に比較している。
		return true
	}

	if targetErr, ok := target.(*Error); ok {
		// Check kind
		if targetErr.Kind != emptyKind && isKindOf(e, targetErr.Kind) {
			return true
		}
	}

	return errors.Is(e.Err, target)
}

func isKindOf(err error, kind ErrorKind) bool {
	return GetKind(err) == kind
}

func GetKind(err error) ErrorKind {
	kind, ok := getKind(err)
	if ok {
		return kind
	} else {
		return InternalError
	}
}

func getKind(err error) (kind ErrorKind, ok bool) {
	if err == nil {
		return emptyKind, false
	}

	// First, pick up the kind from the error itself
	if hasKind, ok := err.(*Error); ok { //nolint: errorlint // errors.As() を使えという指摘だが、ここでは err 変数そのものが ex.Error 型であるかどうかが重要
		if hasKind.Kind != emptyKind {
			return hasKind.Kind, true
		}
	}

	// Now, we check wrapped (nested) errors recursively

	// For the error which wraps single error
	if unwrappable, ok := err.(interface{ Unwrap() error }); ok {
		return getKind(unwrappable.Unwrap())
	}

	// For the error which wraps multiple errors
	if multi, ok := err.(interface{ Unwrap() []error }); ok {
		errs := multi.Unwrap()
		kinds := make([]ErrorKind, 0, len(errs))
		for _, e := range errs {
			if k, ok := getKind(e); ok {
				kinds = append(kinds, k)
			}
		}

		switch len(kinds) {
		case 0:
			// No kind is determined
			return emptyKind, false
		case 1:
			// A kind is determined
			return kinds[0], true
		default:
			// Multiple kinds are found, so we cannot determine the kind.
			// It should be a programming error.
			return emptyKind, false
		}
	}

	return emptyKind, false
}

func WrapAsInvalidArgument(err error, args ...any) error {
	return wrap(err, InvalidArgumentError, args...)
}

func WrapAsUnauthenticated(err error, args ...any) error {
	return wrap(err, UnauthenticatedError, args...)
}

func WrapAsAccessDenied(err error, args ...any) error {
	return wrap(err, AccessDeniedError, args...)
}

func WrapAsNotFound(err error, args ...any) error {
	return wrap(err, NotFoundError, args...)
}

func WrapAsTimeout(err error, args ...any) error {
	return wrap(err, TimeoutError, args...)
}

func WrapAsConflict(err error, args ...any) error {
	return wrap(err, ConflictError, args...)
}

func WrapAsResourceExhausted(err error, args ...any) error {
	return wrap(err, ResourceExhaustedError, args...)
}

func WrapAsCanceled(err error, args ...any) error {
	return wrap(err, CanceledError, args...)
}

func WrapAsDependentService(err error, args ...any) error {
	return wrap(err, DependentServiceError, args...)
}

func WrapAsFailedPrecondition(err error, args ...any) error {
	return wrap(err, FailedPreconditionError, args...)
}
