package user_test

import (
	"testing"
	"user-api/entity/user"

	"github.com/stretchr/testify/assert"
)

var msgErrorInvalidParam string = "Param different from expected when build User"
var msgErrorCreateInvalidUser string = "User created without requested params"

func TestInvalidBuild(t *testing.T) {

	builders := []user.Builder{
		user.NewBuilder().Email(emailTest).Age(ageTest).AddressesIds(addressesIdsTest),
		user.NewBuilder().Name(nameTest).Age(ageTest).AddressesIds(addressesIdsTest),
		user.NewBuilder().Name(nameTest).Email(emailTest).AddressesIds(addressesIdsTest),
	}

	for _, builder := range builders {

		user, err := builder.Build()

		assert.Nil(t, user, msgErrorCreateInvalidUser)
		assert.Error(t, err, msgErrorCreateInvalidUser)
	}
}

func TestValidBuildOnlyRequestParams(t *testing.T) {

	usr, err := user.NewBuilder().Name(nameTest).Email(emailTest).Age(ageTest).Build()
	validateRequestParams(t, usr, err)
}

func TestValidBuildWithAllParams(t *testing.T) {

	usr, err := user.NewBuilder().Id(idTest).Name(nameTest).Email(emailTest).Age(ageTest).AddressesIds(addressesIdsTest).Build()
	assert.Equal(t, usr.ID(), idTest, msgErrorInvalidParam)
	assert.Equal(t, usr.AddressesIDs(), addressesIdsTest, msgErrorInvalidParam)
	validateRequestParams(t, usr, err)
}

func validateRequestParams(t *testing.T, usr user.User, err error) {

	assert.Nil(t, err, "Error occurred when Build a valid new User")
	assert.Equal(t, usr.Name(), nameTest, msgErrorInvalidParam)
	assert.Equal(t, usr.Email(), emailTest, msgErrorInvalidParam)
	assert.Equal(t, usr.Age(), ageTest, msgErrorInvalidParam)
}
