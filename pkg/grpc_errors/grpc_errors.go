package grpc_errors

import (
	"context"
	"database/sql"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
)

var (
	ErrNotFound             = errors.New("Not found")
	ErrNoCtxMetaData        = errors.New("No ctx metadata")
	ErrInvalidDayTimeFormat = errors.New("Invalid day time format")
)

// Parse error and get code
func ParseGRPCErrStatusCode(err error) codes.Code {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return codes.NotFound
	case errors.Is(err, context.Canceled):
		return codes.Canceled
	case errors.Is(err, context.DeadlineExceeded):
		return codes.DeadlineExceeded
	case errors.Is(err, ErrNoCtxMetaData):
		return codes.Unauthenticated
	case errors.Is(err, ErrInvalidDayTimeFormat):
		return codes.InvalidArgument
	case strings.Contains(err.Error(), "no permission"):
		return codes.PermissionDenied
	case strings.Contains(err.Error(), "Validate"):
		return codes.InvalidArgument
	case strings.Contains(err.Error(), "Invalid"):
		return codes.InvalidArgument
	}
	return codes.Internal
}

// Map GRPC errors codes to http status
func MapGRPCErrCodeToHttpStatus(code codes.Code) int {
	switch code {
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.AlreadyExists:
		return http.StatusBadRequest
	case codes.NotFound:
		return http.StatusNotFound
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.Canceled:
		return http.StatusRequestTimeout
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.InvalidArgument:
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}
