package responder

import (
	"encoding/json"
	"net/http"

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
	st, ok := status.FromError(err)
	if !ok {
		Error(w, http.StatusInternalServerError, "internal server error")
		return
	}

	reason := extractErrorReason(st)

	var httpCode int
	var message string
	switch st.Code() {
	case codes.NotFound:
		httpCode = http.StatusNotFound
		message = reason
	case codes.InvalidArgument:
		httpCode = http.StatusBadRequest
		message = reason
	case codes.Unauthenticated:
		httpCode = http.StatusUnauthorized
		message = reason
	case codes.PermissionDenied:
		httpCode = http.StatusForbidden
		message = reason
	case codes.AlreadyExists:
		httpCode = http.StatusConflict
		message = reason
	case codes.FailedPrecondition:
		httpCode = http.StatusBadRequest
		message = reason
	case codes.Internal:
		httpCode = http.StatusInternalServerError
		message = "internal server error"
	case codes.Unavailable:
		httpCode = http.StatusServiceUnavailable
		message = "service unavailable"
	case codes.DeadlineExceeded:
		httpCode = http.StatusGatewayTimeout
		message = "request timeout"
	default:
		httpCode = http.StatusInternalServerError
		message = "internal server error"
	}

	Error(w, httpCode, message)
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
