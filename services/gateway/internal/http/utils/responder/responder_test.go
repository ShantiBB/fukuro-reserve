package responder

import (
	"errors"
	"testing"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGRPCToHTTP_Mapping(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		err         error
		wantCode    int
		wantMessage string
	}{
		{
			name:        "not found",
			err:         grpcErr(codes.NotFound, "user not found"),
			wantCode:    404,
			wantMessage: "user not found",
		},
		{
			name:        "invalid argument",
			err:         grpcErr(codes.InvalidArgument, "validation failed"),
			wantCode:    400,
			wantMessage: "validation failed",
		},
		{
			name:        "unauthenticated",
			err:         grpcErr(codes.Unauthenticated, "invalid token"),
			wantCode:    401,
			wantMessage: "invalid token",
		},
		{
			name:        "permission denied",
			err:         grpcErr(codes.PermissionDenied, "forbidden"),
			wantCode:    403,
			wantMessage: "forbidden",
		},
		{
			name:        "already exists",
			err:         grpcErr(codes.AlreadyExists, "username or email already exists"),
			wantCode:    409,
			wantMessage: "username or email already exists",
		},
		{
			name:        "failed precondition",
			err:         grpcErr(codes.FailedPrecondition, "hotel has rooms"),
			wantCode:    400,
			wantMessage: "hotel has rooms",
		},
		{
			name:        "internal",
			err:         grpcErr(codes.Internal, "internal server error"),
			wantCode:    500,
			wantMessage: "internal server error",
		},
		{
			name:        "unavailable",
			err:         grpcErr(codes.Unavailable, "backend down"),
			wantCode:    503,
			wantMessage: "service unavailable",
		},
		{
			name:        "deadline exceeded",
			err:         grpcErr(codes.DeadlineExceeded, "timeout"),
			wantCode:    504,
			wantMessage: "request timeout",
		},
		{
			name:        "non grpc error",
			err:         errors.New("boom"),
			wantCode:    500,
			wantMessage: "internal server error",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			code, msg, _ := grpcToHTTP(tc.err)
			if code != tc.wantCode {
				t.Fatalf("unexpected code: got=%d want=%d", code, tc.wantCode)
			}
			if msg != tc.wantMessage {
				t.Fatalf("unexpected message: got=%q want=%q", msg, tc.wantMessage)
			}
		})
	}
}

func TestGRPCToHTTP_ValidationDetails(t *testing.T) {
	t.Parallel()

	st := status.New(codes.InvalidArgument, "validation failed")
	stWithDetails, err := st.WithDetails(&errdetails.BadRequest{
		FieldViolations: []*errdetails.BadRequest_FieldViolation{
			{Field: "email", Description: "value must be a valid email address"},
			{Field: "password", Description: "value length must be at least 8 characters"},
		},
	})
	if err != nil {
		t.Fatalf("failed to add details: %v", err)
	}

	code, msg, details := grpcToHTTP(stWithDetails.Err())
	if code != 400 {
		t.Fatalf("unexpected code: got=%d want=400", code)
	}
	if msg != "validation failed" {
		t.Fatalf("unexpected message: got=%q want=%q", msg, "validation failed")
	}
	if len(details) != 2 {
		t.Fatalf("unexpected details len: got=%d want=2", len(details))
	}
	if details[0].Field != "email" || details[1].Field != "password" {
		t.Fatalf("unexpected fields in details: %+v", details)
	}
}

func grpcErr(code codes.Code, reason string) error {
	ei := &errdetails.ErrorInfo{Reason: reason, Domain: "user-service"}
	st, err := status.New(code, "operation failed").WithDetails(ei)
	if err != nil {
		panic(err)
	}
	return st.Err()
}
