package user_test

import (
	"testing"
	"user-api/entity/user"

	"github.com/stretchr/testify/assert"
)

const idTestBuilder string = "123"
const nameTestBuilder string = "Name"
const emailTestBuilder string = "Email@gmail.com"
const ageTestBuilder int = 12

var addressesIdsTestBuilder []string = []string{"1", "2"}

const msgErrorInvalidParamBuilder string = "Param different from expected when build User"
const msgErrorCreateInvalidBuilder string = "User created without requested params"

func TestInvalidBuild(t *testing.T) {

	builders := []user.Builder{
		user.NewBuilder().Email(emailTestBuilder).Age(ageTestBuilder).AddressesIds(addressesIdsTestBuilder),
		user.NewBuilder().Name(nameTestBuilder).Age(ageTestBuilder).AddressesIds(addressesIdsTestBuilder),
		user.NewBuilder().Name(nameTestBuilder).Email(emailTestBuilder).AddressesIds(addressesIdsTestBuilder),
	}

	for _, builder := range builders {

		user, err := builder.Build()

		assert.Nil(t, user, msgErrorInvalidParamBuilder)
		assert.Error(t, err, msgErrorInvalidParamBuilder)
	}
}

func TestValidBuildOnlyRequestParams(t *testing.T) {

	usr, err := user.
		NewBuilder().
		Name(nameTestBuilder).
		Email(emailTestBuilder).
		Age(ageTestBuilder).
		Build()

	validateRequestParams(t, usr, err)
}

func TestValidBuildWithAllParams(t *testing.T) {

	usr, err := user.
		NewBuilder().
		Id(idTestBuilder).
		Name(nameTestBuilder).
		Email(emailTestBuilder).
		Age(ageTestBuilder).
		AddressesIds(addressesIdsTestBuilder).
		Build()

	assert.Equal(t, usr.Id(), idTestBuilder, msgErrorInvalidParamBuilder)
	assert.Equal(t, usr.AddressesIds(), addressesIdsTestBuilder, msgErrorInvalidParamBuilder)
	validateRequestParams(t, usr, err)
}

func validateRequestParams(t *testing.T, usr user.User, err error) {

	assert.Nil(t, err, "Error occurred when Build a valid new User")
	assert.Equal(t, usr.Name(), nameTestBuilder, msgErrorInvalidParamBuilder)
	assert.Equal(t, usr.Email(), emailTestBuilder, msgErrorInvalidParamBuilder)
	assert.Equal(t, usr.Age(), ageTestBuilder, msgErrorInvalidParamBuilder)
}
