package double

import (
	"context"
	"user-api/usecase"
)

type removeUserDouble struct {
	RemoveFunc func(ctx context.Context, request usecase.RemoveUserRequest) error
}

func NewRemoveUserDouble(removeFunc func(ctx context.Context, request usecase.RemoveUserRequest) error) usecase.RemoveUser {
	return &removeUserDouble{removeFunc}
}

func (double *removeUserDouble) Remove(ctx context.Context, request usecase.RemoveUserRequest) error {
	return double.RemoveFunc(ctx, request)
}
