package double

import (
	"context"
	"user-api/usecase"
)

type createUserDouble struct {
	createFunc func(ctx context.Context, request usecase.CreateUserRequest) (usecase.CreateUserResponse, error)
}

func NewCreateUser(create func(ctx context.Context, request usecase.CreateUserRequest) (usecase.CreateUserResponse, error)) usecase.CreateUser {
	return &createUserDouble{createFunc: create}
}

func (create *createUserDouble) Create(ctx context.Context, request usecase.CreateUserRequest) (usecase.CreateUserResponse, error) {
	return create.createFunc(ctx, request)
}

type createUserResponse struct {
	idFunc func() string
}

func NewCreateUserResponse(idFunc func() string) usecase.CreateUserResponse {
	return &createUserResponse{idFunc: idFunc}
}

func (response *createUserResponse) Id() string {
	return response.idFunc()
}
