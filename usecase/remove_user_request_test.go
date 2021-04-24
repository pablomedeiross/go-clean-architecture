package usecase_test

import (
	"testing"
	"user-api/usecase"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewRemoveUserRequest(t *testing.T) {

	request, err := usecase.NewRemoveUserRequest("1")
	assert.NoError(t, err)
	assert.NotNil(t, request)
}

func TestCreateNewRemoveUserRequestWithEmptyId(t *testing.T) {

	request, err := usecase.NewRemoveUserRequest("")
	assert.Nil(t, request)
	assert.Error(t, err)
}

func TestGetName(t *testing.T) {

	expectedName := "1"
	request, err := usecase.NewRemoveUserRequest(expectedName)
	assert.NoError(t, err)
	assert.Equal(t, expectedName, request.Name())
}
