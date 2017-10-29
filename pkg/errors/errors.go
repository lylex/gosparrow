package errors

import (
	"errors"

	"gosparrow/pkg/consts"
)

// APIError represents common api error
type APIError struct {
	Err string `json:"error"`
}

// NewAPIError used to generate a APIError
func NewAPIError(msg string) *APIError {
	return &APIError{
		Err: msg,
	}
}

// Error used to achieve error compliance
func (err *APIError) Error() string {
	return err.Err
}

// Handler related errors
var (
	ErrHandlerBadRequestBody       = errors.New(consts.ErrHandlerBadRequestBodyMsg)
	ErrHandlerUnsupportedMediaType = errors.New(consts.ErrHandlerUnsupportedMediaTypeMsg)
)
