package http

import (
	"context"
	"user-api/usecase"

	"github.com/pkg/errors"
)

const (
	create_controller_without_param_createuser = "CreateUser is a request param for create HttpController"
	create_controller_without_param_removeuser = "RemoveUser is a request param for create HttpController"
	request_param_empty_create_user            = "One of request params are empty to the try create new user"
	request_param_empty_remove_user            = "One of request params are empty to the try remove user"
	couldnt_create_new_user                    = "Couldn't create new user, usecase returning error"
	couldnt_remove_user                        = "Couldn't remove user, usecase returning error"
)

type HttpController interface {
	CreateUser(ctx context.Context, user User) (string, error)
	RemoveUser(ctx context.Context, name string) error
}

type controller struct {
	createUser usecase.CreateUser
	removeUser usecase.RemoveUser
}

func NewHttpController(createUser *usecase.CreateUser, removeUser *usecase.RemoveUser) (HttpController, error) {

	if createUser == nil {
		return nil, errors.New(create_controller_without_param_createuser)
	}

	if removeUser == nil {
		return nil, errors.New(create_controller_without_param_removeuser)
	}

	return &controller{*createUser, *removeUser}, nil
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

func (controller *controller) RemoveUser(ctx context.Context, name string) error {

	req, err := usecase.NewRemoveUserRequest(name)

	if err != nil {
		return errors.Wrap(err, request_param_empty_remove_user)
	}

	err = (controller.removeUser).Remove(ctx, req)

	if err != nil {
		return errors.Wrap(err, couldnt_remove_user)
	}

	return nil
}
