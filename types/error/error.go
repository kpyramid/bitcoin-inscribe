package error

import (
	"context"
	"encoding/json"
	_ "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StatusErrorKey int

type StatusError struct {
	Code    StatusErrorKey
	Message string
}

const (
	StatusErrorUnknown = -1
)

func (a StatusError) Error() string {
	data, _ := json.Marshal(&a)
	return string(data)
}

var _ error = &StatusError{}

func NewStatusError(ctx context.Context, code StatusErrorKey) *StatusError {
	msg, err := GetI18nMessage(ctx, code)
	if err != nil {
		msg = err.Error()
	}

	return &StatusError{
		Code:    code,
		Message: msg,
	}
}

func NewStatusErrorF(ctx context.Context, code StatusErrorKey, args interface{}) *StatusError {
	msg, err := GetI18nMessageF(ctx, code, args)
	if err != nil {
		msg = err.Error()
	}

	return &StatusError{
		Code:    code,
		Message: msg,
	}
}

func NewInternalError(err error) error {
	return status.Error(codes.Internal, err.Error())
}

func AuthErrorFromString(str string) (*StatusError, error) {
	authErr := &StatusError{}
	if err := json.Unmarshal([]byte(str), authErr); err != nil {
		return nil, err
	}

	return authErr, nil
}

func (a StatusError) ProtoError() error {
	return status.Error(codes.InvalidArgument, a.Error())
}
