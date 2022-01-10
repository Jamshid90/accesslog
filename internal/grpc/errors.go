package grpc

import (
	"errors"

	errorpkg "github.com/MedHubUz/access-log/internal/errors"
	"github.com/MedHubUz/access-log/internal/validation"
	epb "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errValidation *validation.ErrValidation
	errNotFound   *errorpkg.ErrNotFound
	errBadRequest *errorpkg.ErrBadRequest
)

func ErrorStatus(err error) *status.Status {
	var (
		st *status.Status
	)

	switch {

	// error bad request
	case errors.As(err, &errBadRequest):
		st = status.New(codes.InvalidArgument, "invalid argument")

	// error validation errors
	case errors.As(err, &errValidation):
		st = status.New(codes.InvalidArgument, "invalid argument")
		br := &epb.BadRequest{}
		for field, des := range errValidation.Errors {
			br.FieldViolations = append(br.FieldViolations, &epb.BadRequest_FieldViolation{
				Field:       field,
				Description: des,
			})
		}
		st, err = st.WithDetails(br)

	// error not found
	case errors.As(err, &errNotFound):
		st = status.New(codes.NotFound, err.Error())

	// error internal
	default:
		st = status.New(codes.Internal, "internal")
		errInfo := &epb.ErrorInfo{
			Reason: err.Error(),
		}
		st, err = st.WithDetails(errInfo)

	}

	return st
}

func Error(err error) error {
	return ErrorStatus(err).Err()
}


