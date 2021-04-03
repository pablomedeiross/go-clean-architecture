package test

import (
	"context"
	"user-api/adapter/http"
)

type controllerDouble struct {
	CreateUserFunc func(ctx context.Context, user http.User) (string, error)
}

func NewUserController(createUserFunc func(ctx context.Context, user http.User) (string, error)) http.HttpController {
	return &controllerDouble{createUserFunc}
}

func (controller *controllerDouble) CreateUser(ctx context.Context, user http.User) (string, error) {
	return controller.CreateUserFunc(ctx, user)
}
