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

// Is checks if the error matches the target error or kind
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

// GetKind returns the kind of the error, or InternalError if no kind is determined
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

// WrapAsInvalidArgument wraps an error as an InvalidArgument error
func WrapAsInvalidArgument(err error, args ...any) error {
	return wrap(err, InvalidArgumentError, args...)
}

// WrapAsUnauthenticated wraps an error as an Unauthenticated error
func WrapAsUnauthenticated(err error, args ...any) error {
	return wrap(err, UnauthenticatedError, args...)
}

// WrapAsAccessDenied wraps an error as an AccessDenied error
func WrapAsAccessDenied(err error, args ...any) error {
	return wrap(err, AccessDeniedError, args...)
}

// WrapAsNotFound wraps an error as a NotFound error
func WrapAsNotFound(err error, args ...any) error {
	return wrap(err, NotFoundError, args...)
}

// WrapAsTimeout wraps an error as a Timeout error
func WrapAsTimeout(err error, args ...any) error {
	return wrap(err, TimeoutError, args...)
}

// WrapAsConflict wraps an error as a Conflict error
func WrapAsConflict(err error, args ...any) error {
	return wrap(err, ConflictError, args...)
}

// WrapAsResourceExhausted wraps an error as a ResourceExhausted error
func WrapAsResourceExhausted(err error, args ...any) error {
	return wrap(err, ResourceExhaustedError, args...)
}

// WrapAsCanceled wraps an error as a Canceled error
func WrapAsCanceled(err error, args ...any) error {
	return wrap(err, CanceledError, args...)
}

// WrapAsDependentService wraps an error as a DependentService error
func WrapAsDependentService(err error, args ...any) error {
	return wrap(err, DependentServiceError, args...)
}

// WrapAsFailedPrecondition wraps an error as a FailedPrecondition error
func WrapAsFailedPrecondition(err error, args ...any) error {
	return wrap(err, FailedPreconditionError, args...)
}
