package usecase_test

import (
	"testing"
	"user-api/usecase"
	"github.com/stretchr/testify/assert"
)

func TestNewCreateUserResponseValidation(t *testing.T) {

	createUserTest, err := usecase.NewCreateUserReponse(0)
	assert.Nil(t, createUserTest)
	assert.EqualError(t, err, "id null to the create CreateUserResponse")
}

func TestID(t *testing.T) {

	idTest := 1
	response, _ := usecase.NewCreateUserReponse(idTest)
	assert.Equal(t, response.ID(), idTest)
}