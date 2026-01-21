package models

import "mini-api-go/server"

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type AppError struct {
	Message string
	Code    int
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(message string, code int) *AppError {
	return &AppError{Message: message, Code: code}
}

func ResponseError(c *server.Context, appError *AppError) {
	err := c.JSON(
		appError.Code,
		ErrorResponse{Error: "error", Message: appError.Message, Code: appError.Code},
	)
	if err != nil {
		return
	}
}
