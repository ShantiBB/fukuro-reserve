package responder

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/consts"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ErrorResponse represents an error response.
type ErrorResponse struct {
	Error string `json:"error"`
}

// JSON sends a JSON response.
func JSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set(consts.HeaderContentType, consts.ContentTypeJSON)
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// Error sends an error response.
func Error(w http.ResponseWriter, code int, message string) {
	JSON(w, code, &ErrorResponse{Error: message})
}

// GRPCError converts gRPC error to HTTP error.
func GRPCError(w http.ResponseWriter, err error) {
	httpCode, message := grpcToHTTP(err)
	Error(w, httpCode, message)
}

// GinJSON sends a JSON response in gin handlers.
func GinJSON(c *gin.Context, code int, payload interface{}) {
	c.JSON(code, payload)
}

// GinError sends an error response in gin handlers.
func GinError(c *gin.Context, code int, message string) {
	c.JSON(code, &ErrorResponse{Error: message})
}

// GinGRPCError converts gRPC error to HTTP error in gin handlers.
func GinGRPCError(c *gin.Context, err error) {
	httpCode, message := grpcToHTTP(err)
	GinError(c, httpCode, message)
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
