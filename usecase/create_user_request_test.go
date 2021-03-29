package usecase_test

import (
	"testing"
	"user-api/usecase"

	"github.com/stretchr/testify/assert"
)

const msgErrorUnexpeted string = "Error to the instantiate CreateUserRequest"

func TestNewCreateUserRequestValidations(t *testing.T) {

	names, emails, ages :=
		[]string{"", "name1", "name2"},
		[]string{"email1@gmail.com", "", "email1@gmail.com"},
		[]int{13, 16, 0}

	for i := range names {
		request, err := usecase.NewCreateUserRequest(names[i], emails[i], ages[i])
		assert.Nil(t, request, "CreateUserRequest created without request params")
		assert.EqualError(t, err, "NewCreateUserRequest called without requested parameter")
	}
}

func TestName(t *testing.T) {

	nameExpected := "Name"
	request, err := usecase.NewCreateUserRequest(nameExpected, "email", 12)
	assert.Nil(t, err, msgErrorUnexpeted)
	assert.Equal(t, request.Name(), nameExpected, "Name different from expected")
}

func TestEmail(t *testing.T) {

	emailExpected := "Name"
	request, err := usecase.NewCreateUserRequest("name", emailExpected, 12)
	assert.Nil(t, err, msgErrorUnexpeted)
	assert.Equal(t, request.Email(), emailExpected, "Email different from expected")
}

func TestAge(t *testing.T) {

	ageExpected := 14
	request, err := usecase.NewCreateUserRequest("name", "email", ageExpected)
	assert.Nil(t, err, msgErrorUnexpeted)
	assert.Equal(t, request.Age(), ageExpected, "Age different from expected")
}
