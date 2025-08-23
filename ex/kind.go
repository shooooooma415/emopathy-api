package ex

import (
	"errors"
)

// ErrorKind represents the type of error
type ErrorKind string

// Error kinds
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
	}
	return InternalError
}

func getKind(err error) (kind ErrorKind, ok bool) {
	if err == nil {
		return emptyKind, false
	}

	if hasKind, ok := err.(*Error); ok { //nolint: errorlint // errors.As() を使えという指摘だが、ここでは err 変数そのものが ex.Error 型であるかどうかが重要
		if hasKind.Kind != emptyKind {
			return hasKind.Kind, true
		}
	}

	if unwrappable, ok := err.(interface{ Unwrap() error }); ok {
		return getKind(unwrappable.Unwrap())
	}

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
			return emptyKind, false
		case 1:
			return kinds[0], true
		default:
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
