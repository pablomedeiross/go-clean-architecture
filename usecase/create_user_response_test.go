package usecase_test

import (
	"testing"
	"user-api/usecase"

	"github.com/stretchr/testify/assert"
)

func TestId(t *testing.T) {

	idTest := "1"
	response, _ := usecase.NewCreateUserReponse(idTest)
	assert.Equal(t, response.Id(), idTest)
}
