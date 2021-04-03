package http

import (
	"context"
	"user-api/usecase"

	"github.com/pkg/errors"
)

const (
	create_controller_without_param = "CreateUser is a request param for create HttpController"
	request_param_empty_create_user = "One of request params are empty to the try create new user"
	couldnt_create_new_user         = "Couldn't create new user, usecase returning error"
)

type HttpController interface {
	CreateUser(ctx context.Context, user User) (string, error)
}

type controller struct {
	createUser usecase.CreateUser
}

func NewHttpController(createUser *usecase.CreateUser) (HttpController, error) {

	if createUser == nil {
		return nil, errors.New(create_controller_without_param)
	}

	return &controller{*createUser}, nil
}

func (controller *controller) CreateUser(ctx context.Context, user User) (string, error) {

	req, err := usecase.NewCreateUserRequest(user.Name, user.Email, user.Age)

	if err != nil {
		return "", errors.Wrap(err, request_param_empty_create_user)
	}

	resp, err := (controller.createUser).Create(ctx, req)

	if err != nil {
		return "", errors.Wrap(err, couldnt_create_new_user)
	}

	return resp.Id(), nil
}
