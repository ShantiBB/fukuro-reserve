package responder

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/consts"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ErrorResponse represents an error response.
type ErrorResponse struct {
	Error   string                 `json:"error"`
	Details []ErrorResponseDetails `json:"details,omitempty"`
}

type ErrorResponseDetails struct {
	Field       string `json:"field,omitempty"`
	Description string `json:"description,omitempty"`
}

// GinGRPCError converts gRPC error to HTTP error in gin handlers.
func GinGRPCError(c *gin.Context, err error) {
	httpCode, message, details := grpcToHTTP(err)
	c.JSON(httpCode, &ErrorResponse{Error: message, Details: details})
}

func grpcToHTTP(err error) (int, string, []ErrorResponseDetails) {
	st, ok := status.FromError(err)
	if !ok {
		return http.StatusInternalServerError, consts.ErrInternalServerError, nil
	}

	reason := extractErrorReason(st)
	details := extractErrorDetails(st)

	switch st.Code() {
	case codes.NotFound:
		return http.StatusNotFound, reason, details
	case codes.InvalidArgument:
		return http.StatusBadRequest, reason, details
	case codes.Unauthenticated:
		return http.StatusUnauthorized, reason, details
	case codes.PermissionDenied:
		return http.StatusForbidden, reason, details
	case codes.AlreadyExists:
		return http.StatusConflict, reason, details
	case codes.FailedPrecondition:
		return http.StatusBadRequest, reason, details
	case codes.Internal:
		return http.StatusInternalServerError, consts.ErrInternalServerError, nil
	case codes.Unavailable:
		return http.StatusServiceUnavailable, consts.ErrServiceUnavailable, nil
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout, consts.ErrRequestTimeout, nil
	default:
		return http.StatusInternalServerError, consts.ErrInternalServerError, nil
	}
}

func extractErrorReason(st *status.Status) string {
	details := st.Details()
	for _, d := range details {
		if errInfo, ok := d.(*errdetails.ErrorInfo); ok {
			if errInfo.Reason != "" {
				return errInfo.Reason
			}
		}
	}

	return st.Message()
}

func extractErrorDetails(st *status.Status) []ErrorResponseDetails {
	details := st.Details()
	if len(details) == 0 {
		return nil
	}

	violations := make([]ErrorResponseDetails, 0)
	for _, d := range details {
		switch typed := d.(type) {
		case *errdetails.BadRequest:
			for _, v := range typed.GetFieldViolations() {
				violations = append(violations, ErrorResponseDetails{
					Field:       v.GetField(),
					Description: v.GetDescription(),
				})
			}
		}
	}

	if len(violations) == 0 {
		return nil
	}

	return violations
}
