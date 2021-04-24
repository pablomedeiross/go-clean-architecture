package double

import (
	"context"
	"user-api/adapter/http"
)

type controllerDouble struct {
	CreateUserFunc func(ctx context.Context, user http.User) (string, error)
	RemoveUserFunc func(ctx context.Context, name string) error
}

func NewUserController(
	createUserFunc func(ctx context.Context, user http.User) (string, error),
	removeUserFunc func(ctx context.Context, name string) error,

) http.HttpController {
	return &controllerDouble{createUserFunc, removeUserFunc}
}

func (controller *controllerDouble) CreateUser(ctx context.Context, user http.User) (string, error) {
	return controller.CreateUserFunc(ctx, user)
}

func (controller *controllerDouble) RemoveUser(ctx context.Context, name string) error {
	return controller.RemoveUserFunc(ctx, name)
}
