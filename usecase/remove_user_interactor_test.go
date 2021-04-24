package usecase_test

import (
	"context"
	"testing"
	"user-api/test/double"
	"user-api/usecase"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestNewRemoveUseInteractor(t *testing.T) {

	respository := double.NewUserRepositoryDouble(nil, nil, nil)
	interactor, err := usecase.NewRemoveUser(*respository)

	assert.NoError(t, err)
	assert.NotNil(t, interactor)
}

func TestNewRemoveUseInteractorWithoutRepository(t *testing.T) {

	interactor, err := usecase.NewRemoveUser(nil)
	assert.Error(t, err)
	assert.Nil(t, interactor)
}

func TestRemoveUser(t *testing.T) {

	repository := double.NewUserRepositoryDouble(
		nil,
		nil,
		func(ctx context.Context, name string) error { return nil },
	)

	interactor, _ := usecase.NewRemoveUser(*repository)

	request, _ := usecase.NewRemoveUserRequest("i")
	err := interactor.Remove(context.Background(), request)

	assert.NoError(t, err)
}

func TestRemoveUserRepositoryReturningError(t *testing.T) {

	repository := double.NewUserRepositoryDouble(
		nil,
		nil,
		func(ctx context.Context, name string) error { return errors.New("Error") },
	)

	interactor, _ := usecase.NewRemoveUser(*repository)

	request, _ := usecase.NewRemoveUserRequest("i")
	err := interactor.Remove(context.Background(), request)

	assert.Error(t, err)
}

func TestRemoveUserWithoutRequestParam(t *testing.T) {

	repository := double.NewUserRepositoryDouble(nil, nil, nil)
	interactor, _ := usecase.NewRemoveUser(*repository)
	err := interactor.Remove(context.Background(), nil)

	assert.Error(t, err)
}
