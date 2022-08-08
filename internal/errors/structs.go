package errors

// FieldViolation is a struct for providing field error details in HTTP error. It matches the same struct in errdetails package.
type FieldViolation struct {
	Field       string
	Description string
}

// Error is an internal errors with stacktrace. It can be converted to a HTTP response.
type Error struct {
	error
	kind            ErrorType
	fieldViolations []FieldViolation
}

// ErrorType defines what the error resulted from.
type ErrorType string

const (
	ErrorReadConfig               ErrorType = "Config read error"
	ErrorParseConfig              ErrorType = "Config parse error"
	ErrorNotImplemented           ErrorType = "Not implement"
	ErrorCanceled                 ErrorType = "Canceled"
	ErrorUnknown                  ErrorType = "Unknown error"
	ErrorInvalidArgument          ErrorType = "Invalid argument"
	ErrorDeadlineExceeded         ErrorType = "Deadline exceeded"
	ErrorNotFound                 ErrorType = "Entity not found"
	ErrorAlreadyExists            ErrorType = "Already exists"
	ErrorInvalidCSV               ErrorType = "CSVFileNotValid"
	ErrorFieldMissing             ErrorType = "FieldNotFound"
	ErrorUnsupportedType          ErrorType = "FieldTypeNotSupported"
	ErrorInvalidParse             ErrorType = "ParseError"
	ErrorIO                       ErrorType = "FileReadError"
	ErrorFailedValidation         ErrorType = "ValidationFailed"
	ErrorInvalidOperation         ErrorType = "OperationFailed"
	ErrorInvalidManifestStructure ErrorType = "ManifestStructureNotValid"
	ErrorInvalidBundle            ErrorType = "BundleNotValid"
	ErrorInvalidPackageManifest   ErrorType = "PackageManifestNotValid"
	ErrorObjectFailedValidation   ErrorType = "ObjectFailedValidation"
)
