package errors

import (
	coreErrors "github.com/hadenlabs/gommon/internal/errors"
)

// New returns an error with the supplied kind and message. If message is empty, a default message,
// for the error kind will be used.
func New(kind coreErrors.ErrorType, msg string) error {
	return coreErrors.New(kind, msg)
}

// Errorf formats according to a format specifier and return an unknown error with the string.
func Errorf(kind coreErrors.ErrorType, format string, args ...interface{}) error {
	return coreErrors.Errorf(kind, format, args...)
}

// Wrap returns an error annotating err with a kind and a stacktrace at the point Wrap is called,
// and the supplied kind and message. If err is nil, Wrap returns nil.
func Wrap(err error, kind coreErrors.ErrorType, msg string) error {
	return coreErrors.Wrap(err, kind, msg)
}

// Wrapf returns an error annotating err with a stack trace at the point Wrapf is called, and the
// kind and format specifier. If err is nil, Wrapf returns nil.
func Wrapf(err error, kind coreErrors.ErrorType, format string, args ...interface{}) error {
	return coreErrors.Wrapf(err, kind, format, args...)
}

// WithFieldViolations returns an error with supplied field
// violations.
func WithFieldViolations(kind coreErrors.ErrorType, msg string, fieldViolations []coreErrors.FieldViolation) error {
	return coreErrors.WithFieldViolations(kind, msg, fieldViolations)
}

// WithValidateError maps a Validate error into an internal error representation.
func WithValidateError(err error) error {
	return coreErrors.WithValidateError(err)
}

// As finds the first error in err's chain that matches target, and if so, sets target to that
// error value and return true.
func As(err error, target interface{}) bool {
	return coreErrors.As(err, target)
}

// Unwrap returns the result of calling the Unwrap method on err, if err's
// type contains an Unwrap method returning error.
// Otherwise, Unwrap returns nil.
//
// Same as Go's errors.Unwrap.
func Unwrap(err error) error {
	return coreErrors.Unwrap(err)
}
