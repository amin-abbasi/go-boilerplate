package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// APIResponse represents a generic API response
type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// SendResponse sends an API response with the given status code, status, message, and optional data or error details
func SendResponse(ctx echo.Context, statusCode int, message string, details ...interface{}) error {
	var responseData, errorDetails interface{}

	if len(details) > 0 {
		if isError(details[0]) {
			errorDetails = details[0]
		} else {
			responseData = details[0]
		}
	}

	response := APIResponse{
		Status:  http.StatusText(statusCode),
		Message: message,
		Data:    responseData,
		Error:   errorDetails,
	}

	return ctx.JSON(statusCode, response)
}

// isError checks if the given value is an error type
func isError(value interface{}) bool {
	_, isError := value.(error)
	return isError
}
