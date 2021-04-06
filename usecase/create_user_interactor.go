package usecase

import (
	"context"
	"user-api/entity/user"

	"github.com/pkg/errors"
)

const (
	msg_repository_nil        string = "The UserRepository is a requested param for create a new CreateUser instance"
	msg_user_already_exists   string = "The user already existis, user's name: "
	msg_error_create_user     string = "There was a error when try to create user in database, user's name: "
	msg_error_create_response string = "Error when try to create response for user named: "
)

type createUserInteractor struct {
	repository UserRepository
}

// NewcreateUser is factory of CreateUser
// *UserRepository is a request param if it don't was passed then a error will returned
func NewCreateUser(repository *UserRepository) (CreateUser, error) {

	if repository == nil {
		return nil, errors.New(msg_repository_nil)
	}

	return &createUserInteractor{*repository}, nil
}

// CreateUser validate if user of same name exists in the repository
// case don't exists then user received is created and persisted in repository
func (interactor *createUserInteractor) Create(ctx context.Context, request CreateUserRequest) (CreateUserResponse, error) {

	var response CreateUserResponse
	var err error

	existingUser, _ := interactor.repository.FindByName(ctx, request.Name())

	if existingUser != nil {
		return nil, errors.New(msg_user_already_exists + request.Name())
	}

	user, _ := createUserFromRequest(request)
	userPersisted, err := interactor.repository.Save(ctx, user)

	if err != nil {
		return nil, errors.Wrap(err, msg_error_create_user+request.Name())
	}

	response, err = newCreateUserReponse(userPersisted.Id())

	if err != nil {
		return nil, errors.Wrap(err, msg_error_create_response+request.Name())
	}

	return response, nil
}

func createUserFromRequest(request CreateUserRequest) (user.User, error) {

	return user.
		NewBuilder().
		Name(request.Name()).
		Email(request.Email()).
		Age(request.Age()).
		Build()
}
