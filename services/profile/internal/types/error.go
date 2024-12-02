package types

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Error represents an application-specific error with a message and code.
// It is designed to encapsulate custom error details for better error handling.
type Error struct {
	Message string // Error message describing the error.
	Code    uint16 // Numeric code representing the type of error.
}

// NewError creates a new custom error with the specified code and message.
// Params:
// - code: The numeric error code.
// - message: A descriptive message for the error.
// Returns: A pointer to the newly created Error instance.
func NewError(code uint16, message string) *Error {
	return &Error{Code: code, Message: message}
}

// NewInternalError creates a new internal server error with a default code (500)
// and the specified message.
// Params:
// - message: A descriptive message for the error.
// Returns: A pointer to the newly created Error instance.
func NewInternalError(message string) *Error {
	return &Error{Code: 500, Message: message}
}

// NewNotFoundError creates a new not found error with a default code (404)
// and the specified message.
// Params:
// - message: A descriptive message for the error.
// Returns: A pointer to the newly created Error instance.
func NewNotFoundError(message string) *Error {
	return &Error{Code: 404, Message: message}
}

// NewPermissionDeniedError creates a new permission denied error with a default code (700)
// and the specified message.
// Params:
// - message: A descriptive message for the error.
// Returns: A pointer to the newly created Error instance.
func NewPermissionDeniedError(message string) *Error {
	return &Error{Code: 700, Message: message}
}

// NewBadRequestError creates a new bad request error with a default code (300)
// and the specified message.
// Params:
// - message: A descriptive message for the error.
// Returns: A pointer to the newly created Error instance.
func NewBadRequestError(message string) *Error {
	return &Error{Code: 300, Message: message}
}

// NewAlreadyExistsError creates a new conflict error indicating that a resource already exists,
// with a default code (409) and the specified message.
// Params:
// - message: A descriptive message for the error.
// Returns: A pointer to the newly created Error instance.
func NewAlreadyExistsError(message string) *Error {
	return &Error{Code: 409, Message: message}
}

// ErrorToGRPCStatus converts the custom error into a gRPC status error based on the error code.
// This ensures interoperability with gRPC error handling.
// Returns: An error in the gRPC status format.
func (c *Error) ErrorToGRPCStatus() error {
	switch c.Code {
	case 500:
		return status.Error(codes.Internal, c.Message) // Maps to gRPC Internal code.
	case 404:
		return status.Error(codes.NotFound, c.Message) // Maps to gRPC NotFound code.
	case 700:
		return status.Error(codes.PermissionDenied, c.Message) // Maps to gRPC PermissionDenied code.
	case 300:
		return status.Error(codes.Aborted, c.Message) // Maps to gRPC Aborted code.
	case 409:
		return status.Error(codes.AlreadyExists, c.Message) // Maps to gRPC AlreadyExists code.
	default:
		return status.Error(codes.Aborted, c.Message) // Default to gRPC Aborted code for unhandled cases.
	}
}
