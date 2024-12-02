package types

import (
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Error represents an application-specific error with a message, code, and success indicator.
type Error struct {
	Message string // A descriptive message for the error.
	Code    uint16 // A numeric code representing the error type.
	Success bool   // Indicates whether the operation was successful (false for errors).
}

// NewError creates a new custom error with the specified code and message.
// Params:
// - code: Numeric code representing the error type.
// - message: A descriptive message for the error.
// Returns: A pointer to the newly created Error instance.
func NewError(code uint16, message string) *Error {
	return &Error{Code: code, Message: message, Success: false}
}

// NewInternalError creates a new internal server error with a predefined code (13)
// and the specified message.
// Params:
// - message: A descriptive message for the error.
// Returns: A pointer to the newly created Error instance.
func NewInternalError(message string) *Error {
	return &Error{Code: 13, Message: message, Success: false}
}

// NewNotFoundError creates a new not found error with a predefined code (5)
// and the specified message.
// Params:
// - message: A descriptive message for the error.
// Returns: A pointer to the newly created Error instance.
func NewNotFoundError(message string) *Error {
	return &Error{Code: 5, Message: message, Success: false}
}

// NewPermissionDeniedError creates a new permission denied error with a predefined code (7)
// and the specified message.
// Params:
// - message: A descriptive message for the error.
// Returns: A pointer to the newly created Error instance.
func NewPermissionDeniedError(message string) *Error {
	return &Error{Code: 7, Message: message, Success: false}
}

// NewBadRequestError creates a new bad request error with a predefined code (10)
// and the specified message.
// Params:
// - message: A descriptive message for the error.
// Returns: A pointer to the newly created Error instance.
func NewBadRequestError(message string) *Error {
	return &Error{Code: 10, Message: message, Success: false}
}

// NewAlreadyExistsError creates a new error indicating a conflict where a resource already exists,
// with a predefined code (6) and the specified message.
// Params:
// - message: A descriptive message for the error.
// Returns: A pointer to the newly created Error instance.
func NewAlreadyExistsError(message string) *Error {
	return &Error{Code: 6, Message: message, Success: false}
}

// ErrorToGRPCStatus converts the custom error to a gRPC status error based on the error code.
func (c *Error) ErrorToHttpStatus() int {

	switch c.Code {
	case 13:
		return fiber.StatusInternalServerError
	case 5:
		return fiber.StatusNotFound
	case 7:
		return fiber.StatusUnauthorized
	case 6:
		return fiber.StatusConflict
	case 10:
		return fiber.StatusBadRequest
	default:
		return fiber.StatusBadRequest
	}
}

func (c *Error) ErrorToJsonMessage() map[string]interface{} {
	return map[string]interface{}{
		"message": c.Message,
		"success": false,
	}
}

func ExtractGRPCErrDetails(err error) *Error {
	details := &Error{}

	st, ok := status.FromError(err)
	if !ok {
		details.Code = uint16(codes.Unknown)
		details.Message = err.Error()
		return details
	}

	details.Code = uint16(st.Code())
	details.Message = st.Message()

	return details
}
