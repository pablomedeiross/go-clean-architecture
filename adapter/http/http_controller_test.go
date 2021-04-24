package http_test

import (
	"context"
	"errors"
	"testing"
	"user-api/adapter/http"
	"user-api/test/double"
	"user-api/usecase"

	"github.com/stretchr/testify/assert"
)

const idTest = "123"

func TestCreateNewHttpController(t *testing.T) {

	uscase := double.NewCreateUser(nil)
	double := double.NewRemoveUserDouble(nil)
	controller, err := http.NewHttpController(&uscase, &double)

	assert.NotNil(t, controller)
	assert.Nil(t, err)
}

func TestErrorWhenCreateNewHttpController(t *testing.T) {

	double := double.NewRemoveUserDouble(nil)
	controller, err := http.NewHttpController(nil, &double)

	assert.Nil(t, controller)
	assert.Error(t, err)
}

func TestCreateNewHttpControllerWithoutRemoveUser(t *testing.T) {

	uscase := double.NewCreateUser(nil)
	controller, err := http.NewHttpController(&uscase, nil)

	assert.Nil(t, controller)
	assert.Error(t, err)
}

func TestCreateNewUserWithSucess(t *testing.T) {

	responseDouble := double.NewCreateUserResponse(
		func() string { return idTest },
	)

	removeCaseDouble := double.NewRemoveUserDouble(nil)
	useCaseDouble := double.NewCreateUser(
		func(ctx context.Context, request usecase.CreateUserRequest) (usecase.CreateUserResponse, error) {
			return responseDouble, nil
		},
	)

	controller, _ := http.NewHttpController(&useCaseDouble, &removeCaseDouble)

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

	removeCaseDouble := double.NewRemoveUserDouble(nil)
	useCaseDouble := double.NewCreateUser(
		func(ctx context.Context, request usecase.CreateUserRequest) (usecase.CreateUserResponse, error) {
			return responseDouble, nil
		},
	)

	controller, _ := http.NewHttpController(&useCaseDouble, &removeCaseDouble)
	idUserCreated, err := controller.CreateUser(context.Background(), http.User{})

	assert.Empty(t, idUserCreated, idTest)
	assert.Error(t, err)
}

func TestCreateUserWithUseCaseReturningError(t *testing.T) {

	removeCaseDouble := double.NewRemoveUserDouble(nil)
	useCaseDouble := double.NewCreateUser(
		func(ctx context.Context, request usecase.CreateUserRequest) (usecase.CreateUserResponse, error) {
			return nil, errors.New("Error")
		},
	)

	controller, _ := http.NewHttpController(&useCaseDouble, &removeCaseDouble)
	idUserCreated, err := controller.CreateUser(context.Background(), http.User{})

	assert.Empty(t, idUserCreated, idTest)
	assert.Error(t, err)
}

func TestRemoveUser(t *testing.T) {

	useCaseDouble := double.NewCreateUser(nil)
	useCaseRemove := double.NewRemoveUserDouble(func(ctx context.Context, request usecase.RemoveUserRequest) error { return nil })

	controller, _ := http.NewHttpController(&useCaseDouble, &useCaseRemove)
	err := controller.RemoveUser(context.Background(), "name")

	assert.NoError(t, err)
}

func TestRemoveUserReturningErrorFromUseCase(t *testing.T) {

	useCaseDouble := double.NewCreateUser(nil)
	useCaseRemove := double.NewRemoveUserDouble(func(ctx context.Context, request usecase.RemoveUserRequest) error { return errors.New("Error") })

	controller, _ := http.NewHttpController(&useCaseDouble, &useCaseRemove)
	err := controller.RemoveUser(context.Background(), "name")

	assert.Error(t, err)
}

func TestRemoveUserReturningErrorFromEmptyName(t *testing.T) {

	useCaseDouble := double.NewCreateUser(nil)
	useCaseRemove := double.NewRemoveUserDouble(func(ctx context.Context, request usecase.RemoveUserRequest) error { return nil })

	controller, _ := http.NewHttpController(&useCaseDouble, &useCaseRemove)
	err := controller.RemoveUser(context.Background(), "")

	assert.Error(t, err)
}
