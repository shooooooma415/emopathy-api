package ex

import (
	"fmt"
)

// New creates a new error with the given message and kind
func New(message string, kind ErrorKind, args ...any) error {
	err := fmt.Errorf(message, args...)
	return wrap(err, kind)
}

// NewInvalidArgument creates a new InvalidArgument error
func NewInvalidArgument(message string, args ...any) error {
	return New(message, InvalidArgumentError, args...)
}

// NewUnauthenticated creates a new Unauthenticated error
func NewUnauthenticated(message string, args ...any) error {
	return New(message, UnauthenticatedError, args...)
}

// NewAccessDenied creates a new AccessDenied error
func NewAccessDenied(message string, args ...any) error {
	return New(message, AccessDeniedError, args...)
}

// NewNotFound creates a new NotFound error
func NewNotFound(message string, args ...any) error {
	return New(message, NotFoundError, args...)
}

// NewTimeout creates a new Timeout error
func NewTimeout(message string, args ...any) error {
	return New(message, TimeoutError, args...)
}

// NewConflict creates a new Conflict error
func NewConflict(message string, args ...any) error {
	return New(message, ConflictError, args...)
}

// NewResourceExhausted creates a new ResourceExhausted error
func NewResourceExhausted(message string, args ...any) error {
	return New(message, ResourceExhaustedError, args...)
}

// NewCanceled creates a new Canceled error
func NewCanceled(message string, args ...any) error {
	return New(message, CanceledError, args...)
}

// NewInternal creates a new Internal error
func NewInternal(message string, args ...any) error {
	return New(message, InternalError, args...)
}

// NewDependentService creates a new DependentService error
func NewDependentService(message string, args ...any) error {
	return New(message, DependentServiceError, args...)
}

// NewFailedPrecondition creates a new FailedPrecondition error
func NewFailedPrecondition(message string, args ...any) error {
	return New(message, FailedPreconditionError, args...)
}

// IsKind checks if the error is of the specified kind
func IsKind(err error, kind ErrorKind) bool {
	return GetKind(err) == kind
}

// IsInvalidArgument checks if the error is an InvalidArgument error
func IsInvalidArgument(err error) bool {
	return IsKind(err, InvalidArgumentError)
}

// IsUnauthenticated checks if the error is an Unauthenticated error
func IsUnauthenticated(err error) bool {
	return IsKind(err, UnauthenticatedError)
}

// IsAccessDenied checks if the error is an AccessDenied error
func IsAccessDenied(err error) bool {
	return IsKind(err, AccessDeniedError)
}

// IsNotFound checks if the error is a NotFound error
func IsNotFound(err error) bool {
	return IsKind(err, NotFoundError)
}

// IsTimeout checks if the error is a Timeout error
func IsTimeout(err error) bool {
	return IsKind(err, TimeoutError)
}

// IsConflict checks if the error is a Conflict error
func IsConflict(err error) bool {
	return IsKind(err, ConflictError)
}

// IsResourceExhausted checks if the error is a ResourceExhausted error
func IsResourceExhausted(err error) bool {
	return IsKind(err, ResourceExhaustedError)
}

// IsCanceled checks if the error is a Canceled error
func IsCanceled(err error) bool {
	return IsKind(err, CanceledError)
}

// IsInternal checks if the error is an Internal error
func IsInternal(err error) bool {
	return IsKind(err, InternalError)
}

// IsDependentService checks if the error is a DependentService error
func IsDependentService(err error) bool {
	return IsKind(err, DependentServiceError)
}

// IsFailedPrecondition checks if the error is a FailedPrecondition error
func IsFailedPrecondition(err error) bool {
	return IsKind(err, FailedPreconditionError)
}

// GetArgs returns the arguments associated with the error
func GetArgs(err error) []any {
	if exErr, ok := err.(*Error); ok {
		return exErr.Args
	}
	return nil
}

// GetStack returns the stack trace associated with the error
func GetStack(err error) []string {
	if exErr, ok := err.(*Error); ok {
		return exErr.Stack
	}
	return nil
}

// WithArgs adds additional arguments to an existing error
func WithArgs(err error, args ...any) error {
	if err == nil {
		return nil
	}

	if exErr, ok := err.(*Error); ok {
		exErr.Args = append(exErr.Args, args...)
		return exErr
	}

	// If it's not an ex.Error, wrap it
	return Wrap(err, args...)
}
