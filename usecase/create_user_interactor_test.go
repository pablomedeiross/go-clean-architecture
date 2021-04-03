package usecase_test

import (
	"context"
	"testing"
	"user-api/entity/user"
	"user-api/usecase"
	"user-api/usecase/double"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

const (
	id_user_creation    string = "1234"
	name_user_creation  string = "Name"
	email_user_creation string = "Email"
	age_user_creation   int    = 42
)

var userExpected user.User
var request usecase.CreateUserRequest

func init() {

	request, _ = usecase.
		NewCreateUserRequest(name_user_creation, email_user_creation, age_user_creation)

	userExpected, _ = user.
		NewBuilder().
		Id("123455").
		Name("name").
		Email("email").
		Age(123).
		Build()
}

func TestUserCreationWithSucess(t *testing.T) {

	var repository *usecase.UserRepository = double.NewUserRepositoryDouble(

		func(ctx context.Context, user user.User) (user.User, error) {
			return userExpected, nil
		},

		func(ctx context.Context, name string) (user.User, error) {
			return nil, usecase.NewUserDontExistError("")
		},
	)

	useCase, _ := usecase.NewCreateUser(repository)

	response, err := useCase.Create(context.Background(), request)
	assert.Nil(t, err, "Unexpected error to the create user")
	assert.Equal(t, response.Id(), userExpected.Id(), "Invalid Id from userCreation")
}

func TestUserCreationAlreadyCreated(t *testing.T) {

	var repository *usecase.UserRepository = double.NewUserRepositoryDouble(
		nil,
		func(ctx context.Context, name string) (user.User, error) { return userExpected, nil },
	)

	useCase, _ := usecase.NewCreateUser(repository)

	response, err := useCase.Create(context.Background(), request)
	assert.Error(t, err, "User created even that another user have already created")
	assert.Nil(t, response, "Response created even that another user exist in repository")
}

func TestCreateWithRepositoryReturningErrorToTryCreateUser(t *testing.T) {

	var repository *usecase.UserRepository = double.NewUserRepositoryDouble(
		func(ctx context.Context, user user.User) (user.User, error) { return nil, errors.New("Error") },
		func(ctx context.Context, name string) (user.User, error) { return nil, nil },
	)

	useCase, _ := usecase.NewCreateUser(repository)

	response, err := useCase.Create(context.Background(), request)
	assert.Error(t, err, "Create method working without requested param")
	assert.Nil(t, response, "Reponse create even without requested param")
}

func TestCreateErrorWhenCreateUserWithRepositoryReturn(t *testing.T) {

	usrReturnedFromRepo, _ := user.
		NewBuilder().
		Name("name").
		Email("email").
		Age(123).
		Build()

	var repository *usecase.UserRepository = double.NewUserRepositoryDouble(
		func(ctx context.Context, user user.User) (user.User, error) { return usrReturnedFromRepo, nil },
		func(ctx context.Context, name string) (user.User, error) { return nil, nil },
	)

	useCase, _ := usecase.NewCreateUser(repository)

	response, err := useCase.Create(context.Background(), request)
	assert.Error(t, err, "Repository returning User without id after creation")
	assert.Nil(t, response, "Reponse create even without requested param Id")
}
