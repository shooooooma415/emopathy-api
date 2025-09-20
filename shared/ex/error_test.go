package ex

import (
	"errors"
	"testing"
)

func TestError_Wrap(t *testing.T) {
	originalErr := errors.New("original error")
	wrappedErr := Wrap(originalErr, "additional context")

	if wrappedErr == nil {
		t.Fatal("wrapped error should not be nil")
	}

	exErr, ok := wrappedErr.(*Error)
	if !ok {
		t.Fatal("wrapped error should be of type *Error")
	}

	if exErr.Err != originalErr {
		t.Errorf("expected original error, got %v", exErr.Err)
	}

	if len(exErr.Args) != 1 {
		t.Errorf("expected 1 arg, got %d", len(exErr.Args))
	}

	if exErr.Args[0] != "additional context" {
		t.Errorf("expected 'additional context', got %v", exErr.Args[0])
	}
}

func TestError_WrapAsNotFound(t *testing.T) {
	originalErr := errors.New("user not found")
	wrappedErr := WrapAsNotFound(originalErr, "user_id", "123")

	if wrappedErr == nil {
		t.Fatal("wrapped error should not be nil")
	}

	exErr, ok := wrappedErr.(*Error)
	if !ok {
		t.Fatal("wrapped error should be of type *Error")
	}

	if exErr.Kind != NotFoundError {
		t.Errorf("expected NotFoundError, got %s", exErr.Kind)
	}

	if len(exErr.Args) != 2 {
		t.Errorf("expected 2 args, got %d", len(exErr.Args))
	}
}

func TestError_GetKind(t *testing.T) {
	originalErr := errors.New("not found")
	wrappedErr := WrapAsNotFound(originalErr)

	kind := GetKind(wrappedErr)
	if kind != NotFoundError {
		t.Errorf("expected NotFoundError, got %s", kind)
	}
}

func TestError_IsNotFound(t *testing.T) {
	originalErr := errors.New("not found")
	wrappedErr := WrapAsNotFound(originalErr)

	if !IsNotFound(wrappedErr) {
		t.Error("error should be identified as NotFound")
	}

	if IsNotFound(originalErr) {
		t.Error("original error should not be identified as NotFound")
	}
}

func TestError_NewNotFound(t *testing.T) {
	err := NewNotFound("user %s not found", "123")

	if err == nil {
		t.Fatal("error should not be nil")
	}

	if !IsNotFound(err) {
		t.Error("error should be identified as NotFound")
	}

	expectedMsg := "user 123 not found"
	if err.Error() != expectedMsg {
		t.Errorf("expected '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestError_Recover(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("should not panic")
		}
	}()

	err := Recover(nil)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
}

func TestError_RecoverWithError(t *testing.T) {
	originalErr := errors.New("panic error")

	err := Recover(originalErr)
	if err == nil {
		t.Fatal("error should not be nil")
	}

	if !IsInternal(err) {
		t.Error("recovered error should be Internal")
	}
}

func TestError_RecoverWithNonError(t *testing.T) {
	panicValue := "panic string"

	err := Recover(panicValue)
	if err == nil {
		t.Fatal("error should not be nil")
	}

	if !IsInternal(err) {
		t.Error("recovered error should be Internal")
	}
}
