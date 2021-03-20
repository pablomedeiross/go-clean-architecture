package user_test

import (
	"testing"
	"user-api/entity/user"

	"github.com/stretchr/testify/assert"
)

const idTestUser string = "123"
const nameTestUser string = "Name"
const emailTestUser string = "Email@gmail.com"
const ageTestUser int = 12

var addressesIdsTestUser []string = []string{"1", "2"}

const msgErrorInvalidParamUser string = "Param different from expected when build User"
const msgErrorCreateInvalidUser string = "User created without requested params"

var userTest user.User

func init() {

	userTest, _ = user.
		NewBuilder().
		Id(idTestUser).
		Name(nameTestUser).
		Email(emailTestUser).
		Age(ageTestUser).
		AddressesIds(addressesIdsTestUser).
		Build()
}

func TestGetId(t *testing.T) {

	assert.Equal(t, idTestUser, userTest.Id(), "Invalid id return from Id()")
}

func TestGetName(t *testing.T) {

	assert.Equal(t, nameTestUser, userTest.Name(), "Invalid name returned from Name()")
}

func TestGetEmail(t *testing.T) {

	assert.Equal(t, emailTestUser, userTest.Email(), "Invalid email returned from Email()")
}

func TestGetAge(t *testing.T) {

	assert.Equal(t, ageTestUser, userTest.Age(), "Invalid age returned from Age()")
}

func TestGetAdressesIDs(t *testing.T) {

	assert.Equal(t, addressesIdsTestUser, userTest.AddressesIds(), "Invalid addresses returned from AddressesIDs()")
}

func TestAddAddressInUserWithSameAddress(t *testing.T) {

	addressId := addressesIdsTestUser[0]
	err := userTest.AddAddressId(addressId)

	msgErrorExpected := "AddressId already exists in User: " + userTest.Name()
	assert.EqualError(t, err, msgErrorExpected, "Error AddAddress add same address in User")
}

func TestAddAddressIDInUserWithoutAddressId(t *testing.T) {

	addressID := "56789"
	usr, _ := user.
		NewBuilder().
		Id(idTestUser).
		Name(nameTestUser).
		Email(emailTestUser).
		Age(ageTestUser).
		Build()

	err := usr.AddAddressId(addressID)

	assert.Equal(t, usr.AddressesIds()[0], addressID, "AddressID is different from expected")
	assert.Nil(t, err, "Error occurred when Address was instantiated")
}
