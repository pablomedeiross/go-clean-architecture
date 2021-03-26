package usecase

import (
	"user-api/entity/user"

	"github.com/pkg/errors"
)

const (
	msg_repository_nil        string = "The UserRepository is a requested param for create a new CreateUser instance"
	msg_user_already_exists   string = "The user already existis, user's name: "
	msg_error_create_user     string = "There was a error when try to create user in database, user's name: "
	msg_error_create_response string = "Error when try to create response for user named: "
)

// CreateUser is a usecase for creation of new user in application
type CreateUser interface {
	Create(request CreateUserRequest) (CreateUserResponse, error)
}

type createUserInteractor struct {
	repository *UserRepository
}

func NewCreateUser(repository UserRepository) (CreateUser, error) {

	if repository == nil {
		return nil, errors.New(msg_repository_nil)
	}

	return &createUserInteractor{&repository}, nil
}

func (interactor *createUserInteractor) Create(request CreateUserRequest) (CreateUserResponse, error) {

	var response CreateUserResponse
	var err error

	existingUser := (*interactor.repository).FindByName(request.Name())

	if existingUser != nil {
		return nil, errors.New(msg_user_already_exists + request.Name())
	}

	user := createUserFromRequest(request)
	persistedUser, err := (*interactor.repository).Save(user)

	if err != nil {
		return nil, errors.Wrap(err, msg_error_create_user+request.Name())
	}

	response, err = newCreateUserReponse(persistedUser.Id())

	if err != nil {
		return nil, errors.Wrap(err, msg_error_create_response+request.Name())
	}

	return response, nil
}

func createUserFromRequest(request CreateUserRequest) user.User {

	usr, _ := user.
		NewBuilder().
		Name(request.Name()).
		Email(request.Email()).
		Age(request.Age()).
		Build()

	return usr
}
