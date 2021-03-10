package usecase_test

import (
	"testing"
	"user-api/usecase"
	"github.com/stretchr/testify/assert"
)

func TestNewCreateUserRequestValidations(t *testing.T) {

	names, emails, ages := 
		[]string{"", "name1", "name2"},
		[]string{"email1@gmail.com", "", "email1@gmail.com"},
		[]int{13,16,0}


	request, err := usecase.NewCreateUserRequest(names[i],emails[i], ages[i])
	assert.EqualError(t, err, "NewCreateUserRequest called without requested parameter")
}

// func TestName(t *testing.T) {} 
