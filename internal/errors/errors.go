package errors

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

// Kind returns the error kind.
func (e *Error) Kind() ErrorType {
	return e.kind
}

// FieldViolations returns a structure that represents field validation errors.
func (e *Error) FieldViolations() []FieldViolation {
	return e.fieldViolations
}

// New returns an error with the supplied kind and message. If message is empty, a default message,
// for the error kind will be used.
func New(kind ErrorType, msg string) error {
	if msg == "" {
		msg = string(kind)
	}
	return &Error{
		error: errors.New(msg),
		kind:  kind,
	}
}

// Errorf formats according to a format specifier and return an unknown error with the string.
func Errorf(kind ErrorType, format string, args ...interface{}) error {
	return New(kind, fmt.Sprintf(format, args...))
}

// Wrap returns an error annotating err with a kind and a stacktrace at the point Wrap is called,
// and the supplied kind and message. If err is nil, Wrap returns nil.
func Wrap(err error, kind ErrorType, msg string) error {
	if err == nil {
		return nil
	}
	if msg == "" {
		msg = string(kind)
	}
	return &Error{
		error: errors.Wrap(err, msg),
		kind:  kind,
	}
}

// Wrapf returns an error annotating err with a stack trace at the point Wrapf is called, and the
// kind and format specifier. If err is nil, Wrapf returns nil.
func Wrapf(err error, kind ErrorType, format string, args ...interface{}) error {
	return Wrap(err, kind, fmt.Sprintf(format, args...))
}

// WithFieldViolations returns an error with supplied field
// violations.
func WithFieldViolations(kind ErrorType, msg string, fieldViolations []FieldViolation) error {
	if msg == "" {
		msg = string(kind)
	}
	return &Error{
		error:           errors.New(msg),
		kind:            kind,
		fieldViolations: fieldViolations,
	}
}

// WithValidateError maps a Validate error into an internal error representation.
func WithValidateError(err error) error {
	if err == nil {
		return nil
	}
	errWithType := err.(validator.ValidationErrors) //nolint:errorlint
	if len(errWithType) == 0 {
		return Wrap(err, ErrorUnknown, "")
	}
	fieldViolations := make([]FieldViolation, 0, len(errWithType))
	for _, fieldError := range errWithType {
		fieldViolations = append(fieldViolations, FieldViolation{
			Field:       ToUnderScore(fieldError.Field()),
			Description: fieldError.Tag(),
		})
	}
	// Assume any field violations corresponds to error returns from validator, which is invalid argument.
	return WithFieldViolations(ErrorInvalidArgument, "", fieldViolations)
}

// IsKind checks whether any error in err's chain matches the error kind.
func IsKind(err error, kind ErrorType) bool {
	ie := &Error{}
	if As(err, &ie) {
		return ie.kind == kind
	}
	return false
}

// As finds the first error in err's chain that matches target, and if so, sets target to that
// error value and return true.
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

// Unwrap returns the result of calling the Unwrap method on err, if err's
// type contains an Unwrap method returning error.
// Otherwise, Unwrap returns nil.
//
// Same as Go's errors.Unwrap.
func Unwrap(err error) error {
	return errors.Unwrap(err)
}
