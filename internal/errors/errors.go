package errors

import (
	"fmt"
	"net/http"
)

// Error represents a custom error with context
type Error struct {
	Type    ErrorType `json:"type"`
	Code    string    `json:"code"`
	Message string    `json:"message"`
	Details string    `json:"details,omitempty"`
	Cause   error     `json:"-"`
}

// ErrorType represents the category of error
type ErrorType string

const (
	// ErrorTypeValidation represents validation errors
	ErrorTypeValidation ErrorType = "validation"
	// ErrorTypeAuthentication represents authentication errors
	ErrorTypeAuthentication ErrorType = "authentication"
	// ErrorTypeAuthorization represents authorization errors
	ErrorTypeAuthorization ErrorType = "authorization"
	// ErrorTypeNetwork represents network/connectivity errors
	ErrorTypeNetwork ErrorType = "network"
	// ErrorTypeInternal represents internal server errors
	ErrorTypeInternal ErrorType = "internal"
	// ErrorTypeExternal represents external service errors
	ErrorTypeExternal ErrorType = "external"
	// ErrorTypeConfiguration represents configuration errors
	ErrorTypeConfiguration ErrorType = "configuration"
	// ErrorTypeSignature represents signature/crypto errors
	ErrorTypeSignature ErrorType = "signature"
)

// Error implements the error interface
func (e *Error) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %s (caused by: %v)", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// Unwrap returns the underlying error
func (e *Error) Unwrap() error {
	return e.Cause
}

// New creates a new custom error
func New(errorType ErrorType, code, message string) *Error {
	return &Error{
		Type:    errorType,
		Code:    code,
		Message: message,
	}
}

// Wrap wraps an existing error with context
func Wrap(err error, errorType ErrorType, code, message string) *Error {
	return &Error{
		Type:    errorType,
		Code:    code,
		Message: message,
		Cause:   err,
	}
}

// WithDetails adds details to an error
func (e *Error) WithDetails(details string) *Error {
	e.Details = details
	return e
}

// GetHTTPStatus returns appropriate HTTP status code for the error type
func (e *Error) GetHTTPStatus() int {
	switch e.Type {
	case ErrorTypeValidation:
		return http.StatusBadRequest
	case ErrorTypeAuthentication:
		return http.StatusUnauthorized
	case ErrorTypeAuthorization:
		return http.StatusForbidden
	case ErrorTypeNetwork:
		return http.StatusServiceUnavailable
	case ErrorTypeExternal:
		return http.StatusBadGateway
	case ErrorTypeConfiguration:
		return http.StatusInternalServerError
	case ErrorTypeSignature:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

// Predefined error constructors for common cases

// NewValidationError creates a validation error
func NewValidationError(message string) *Error {
	return New(ErrorTypeValidation, "VALIDATION_ERROR", message)
}

// NewAuthenticationError creates an authentication error
func NewAuthenticationError(message string) *Error {
	return New(ErrorTypeAuthentication, "AUTHENTICATION_ERROR", message)
}

// NewAuthorizationError creates an authorization error
func NewAuthorizationError(message string) *Error {
	return New(ErrorTypeAuthorization, "AUTHORIZATION_ERROR", message)
}

// NewNetworkError creates a network error
func NewNetworkError(err error, message string) *Error {
	return Wrap(err, ErrorTypeNetwork, "NETWORK_ERROR", message)
}

// NewInternalError creates an internal error
func NewInternalError(err error, message string) *Error {
	return Wrap(err, ErrorTypeInternal, "INTERNAL_ERROR", message)
}

// NewExternalError creates an external service error
func NewExternalError(err error, message string) *Error {
	return Wrap(err, ErrorTypeExternal, "EXTERNAL_ERROR", message)
}

// NewConfigurationError creates a configuration error
func NewConfigurationError(message string) *Error {
	return New(ErrorTypeConfiguration, "CONFIGURATION_ERROR", message)
}

// NewSignatureError creates a signature error
func NewSignatureError(err error, message string) *Error {
	return Wrap(err, ErrorTypeSignature, "SIGNATURE_ERROR", message)
}

// ErrorResponse represents an error response structure for APIs
type ErrorResponse struct {
	Error     *Error `json:"error"`
	RequestID string `json:"request_id,omitempty"`
	Timestamp string `json:"timestamp"`
	Path      string `json:"path,omitempty"`
}

// NewErrorResponse creates a new error response
func NewErrorResponse(err *Error, requestID, path string) *ErrorResponse {
	return &ErrorResponse{
		Error:     err,
		RequestID: requestID,
		Timestamp: fmt.Sprintf("%d", 0), // Will be set by middleware
		Path:      path,
	}
}
