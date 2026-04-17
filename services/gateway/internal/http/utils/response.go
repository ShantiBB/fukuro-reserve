package utils

import (
	"encoding/json"
	"net/http"
	"strconv"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// RespondJSON sends a JSON response
func RespondJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// RespondError sends an error response
func RespondError(w http.ResponseWriter, code int, message string) {
	RespondJSON(w, code, &ErrorResponse{Error: message})
}

// RespondGRPCError converts gRPC error to HTTP error
func RespondGRPCError(w http.ResponseWriter, err error) {
	st, ok := status.FromError(err)
	if !ok {
		RespondError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	// Try to extract the error reason from ErrorInfo details
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

	RespondError(w, httpCode, message)
}

// extractErrorReason extracts the specific error reason from gRPC error details
func extractErrorReason(st *status.Status) string {
	// Get the ErrorInfo details if available
	details := st.Details()
	for _, d := range details {
		if errInfo, ok := d.(*errdetails.ErrorInfo); ok {
			if errInfo.Reason != "" {
				return errInfo.Reason
			}
		}
	}
	// Fallback to the default message if no reason found
	return st.Message()
}

// ParseInt64 parses a string to int64
func ParseInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

// ParseUint32 parses a string to uint32
func ParseUint32(s string) uint32 {
	i, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0
	}
	return uint32(i)
}

// ParseUint64 parses a string to uint64
func ParseUint64(s string) uint64 {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	return i
}
