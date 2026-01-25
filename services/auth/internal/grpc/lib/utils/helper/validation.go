package helper

import (
	"errors"

	"buf.build/go/protovalidate"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleValidationErr(err error) error {
	if err == nil {
		return nil
	}

	var valErr *protovalidate.ValidationError
	if !errors.As(err, &valErr) {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	br := &errdetails.BadRequest{}
	for _, v := range valErr.ToProto().GetViolations() {
		var field string
		if elements := v.GetField().GetElements(); len(elements) > 0 {
			field = elements[len(elements)-1].GetFieldName()
		}

		br.FieldViolations = append(
			br.FieldViolations, &errdetails.BadRequest_FieldViolation{
				Field:       field,
				Description: v.GetMessage(),
			},
		)
	}

	st, _ := status.New(codes.InvalidArgument, "validation failed").WithDetails(br)
	return st.Err()
}
