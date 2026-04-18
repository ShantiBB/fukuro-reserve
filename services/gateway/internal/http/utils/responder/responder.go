package responder

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ErrorResponse represents an error response.
type ErrorResponse struct {
	Error string `json:"error"`
}

// GinGRPCError converts gRPC error to HTTP error in gin handlers.
func GinGRPCError(c *gin.Context, err error) {
	httpCode, message := grpcToHTTP(err)
	c.JSON(httpCode, &ErrorResponse{Error: message})
}

func grpcToHTTP(err error) (int, string) {
	st, ok := status.FromError(err)
	if !ok {
		return http.StatusInternalServerError, "internal server error"
	}

	reason := extractErrorReason(st)

	switch st.Code() {
	case codes.NotFound:
		return http.StatusNotFound, reason
	case codes.InvalidArgument:
		return http.StatusBadRequest, reason
	case codes.Unauthenticated:
		return http.StatusUnauthorized, reason
	case codes.PermissionDenied:
		return http.StatusForbidden, reason
	case codes.AlreadyExists:
		return http.StatusConflict, reason
	case codes.FailedPrecondition:
		return http.StatusBadRequest, reason
	case codes.Internal:
		return http.StatusInternalServerError, "internal server error"
	case codes.Unavailable:
		return http.StatusServiceUnavailable, "service unavailable"
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout, "request timeout"
	default:
		return http.StatusInternalServerError, "internal server error"
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
