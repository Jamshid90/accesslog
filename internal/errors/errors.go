package errors

import (
	"errors"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v4"
)

var (
	BadRequest    = errors.New(GetHTTPStatusText(http.StatusBadRequest))
	ErrorNotFound = NewErrNotFound("object")

	// database
	DBErrNoRows = pgx.ErrNoRows
)

// GetHTTPStatusText ...
func GetHTTPStatusText(statusCode int) string {
	return strings.ToLower(http.StatusText(statusCode))
}

// ErrNotFound ...
func NewErrNotFound(text string) *ErrNotFound {
	return &ErrNotFound{text}
}

type ErrNotFound struct {
	name string
}

func (e *ErrNotFound) Error() string {
	return e.name + " not found"
}

// ErrBadRequest ...
func NewErrBadRequest(err error, message string) *ErrBadRequest {
	return &ErrBadRequest{err, message}
}

type ErrBadRequest struct {
	Err     error
	Message string
}

func (e ErrBadRequest) Error() string {
	return e.Message
}
