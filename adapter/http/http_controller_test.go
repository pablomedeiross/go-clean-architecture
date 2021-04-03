package http_test

import (
	"context"
	"errors"
	"testing"
	"user-api/adapter/http"
	"user-api/adapter/http/double"
	"user-api/usecase"

	"github.com/stretchr/testify/assert"
)

const idTest = "123"

func TestCreateNewHttpController(t *testing.T) {

	uscase := double.NewCreateUser(nil)
	controller, err := http.NewHttpController(&uscase)

	assert.NotNil(t, controller)
	assert.Nil(t, err)
}

func TestErrorWhenCreateNewHttpController(t *testing.T) {

	controller, err := http.NewHttpController(nil)

	assert.Nil(t, controller)
	assert.Error(t, err)
}

func TestCreateNewUserWithSucess(t *testing.T) {

	responseDouble := double.NewCreateUserResponse(
		func() string { return idTest },
	)

	useCaseDouble := double.NewCreateUser(
		func(ctx context.Context, request usecase.CreateUserRequest) (usecase.CreateUserResponse, error) {
			return responseDouble, nil
		},
	)

	controller, _ := http.NewHttpController(&useCaseDouble)

	idUserCreated, err := controller.CreateUser(

		context.Background(),

		http.User{
			Name:  "name",
			Email: "email@mail.com",
			Age:   12,
		},
	)

	assert.Equal(t, idUserCreated, idTest)
	assert.Nil(t, err)
}

func TestCreateUserWithRequestParamIsZero(t *testing.T) {

	responseDouble := double.NewCreateUserResponse(
		func() string { return idTest },
	)

	useCaseDouble := double.NewCreateUser(
		func(ctx context.Context, request usecase.CreateUserRequest) (usecase.CreateUserResponse, error) {
			return responseDouble, nil
		},
	)

	controller, _ := http.NewHttpController(&useCaseDouble)
	idUserCreated, err := controller.CreateUser(context.Background(), http.User{})

	assert.Empty(t, idUserCreated, idTest)
	assert.Error(t, err)
}

func TestCreateUserWithUseCaseReturningError(t *testing.T) {

	useCaseDouble := double.NewCreateUser(
		func(ctx context.Context, request usecase.CreateUserRequest) (usecase.CreateUserResponse, error) {
			return nil, errors.New("Error")
		},
	)

	controller, _ := http.NewHttpController(&useCaseDouble)
	idUserCreated, err := controller.CreateUser(context.Background(), http.User{})

	assert.Empty(t, idUserCreated, idTest)
	assert.Error(t, err)
}
