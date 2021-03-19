package user_test

import (
	"testing"
	"user-api/entity/user"

	"github.com/stretchr/testify/assert"
)

func TestInvalidBuild(t *testing.T) {

	builders := []user.Builder{
		user.NewBuilder().Email(emailTest).Age(ageTest).AddressesIds(addressesIdsTest),
		user.NewBuilder().Name(nameTest).Age(ageTest).AddressesIds(addressesIdsTest),
		user.NewBuilder().Name(nameTest).Email(emailTest).AddressesIds(addressesIdsTest),
	}

	for _, builder := range builders {

		user, err := builder.Build()

		assert.Nil(t, user)
		assert.Error(t, err)
	}
}

func TestValidBuildOnlyRequestParams(t *testing.T) {

	usr, err := user.NewBuilder().Name(nameTest).Email(emailTest).Age(ageTest).Build()
	assert.Nil(t, err)
	validateRequestParams(t, usr)
}

func TestValidBuildWithAllParams(t *testing.T) {

	usr, err := user.NewBuilder().Id(idTest).Name(nameTest).Email(emailTest).Age(ageTest).AddressesIds(addressesIdsTest).Build()
	assert.Nil(t, err)
	assert.Equal(t, usr.ID(), idTest)
	assert.Equal(t, usr.AddressesIDs(), addressesIdsTest)
	validateRequestParams(t, usr)
}

func validateRequestParams(t *testing.T, usr user.User) {

	assert.Equal(t, usr.Name(), nameTest)
	assert.Equal(t, usr.Email(), emailTest)
	assert.Equal(t, usr.Age(), ageTest)
}
