package user_test

import (
	"testing"
	"user-api/entity/user"

	"github.com/stretchr/testify/assert"
)

var userTest user.User

func init() {

	userTest, _ = user.
		NewBuilder().
		Id(idTest).
		Name(nameTest).
		Email(emailTest).
		Age(ageTest).
		AddressesIds(addressesIdsTest).
		Build()
}

func TestGetName(t *testing.T) {

	nameExpected := nameTest
	assert.Equal(t, nameExpected, userTest.Name(), "Invalid name returned from Name()")
}

func TestGetEmail(t *testing.T) {

	emailExpected := emailTest
	assert.Equal(t, emailExpected, userTest.Email(), "Invalid email returned from Email()")
}

func TestGetAge(t *testing.T) {

	ageExpected := ageTest
	assert.Equal(t, ageExpected, userTest.Age(), "Invalid age returned from Age()")
}

func TestGetAdressesIDs(t *testing.T) {

	addressesIDsExpected := addressesIdsTest
	assert.Equal(t, addressesIDsExpected, userTest.AddressesIDs(), "Invalid addresses returned from AddressesIDs()")
}

func TestAddAddressInUserWithSameAddress(t *testing.T) {

	addressId := addressesIdsTest[0]
	err := userTest.AddAddressID(addressId)

	msgErrorExpected := "AddressId already exists in User: " + userTest.Name()
	assert.EqualError(t, err, msgErrorExpected, "Error AddAddress add same address in User")
}

func TestAddAddressIDInUserWithoutAddressId(t *testing.T) {

	addressID := "56789"
	usr, _ := user.
		NewBuilder().
		Id(idTest).
		Name(nameTest).
		Email(emailTest).
		Age(ageTest).
		Build()

	err := usr.AddAddressID(addressID)

	assert.Equal(t, usr.AddressesIDs()[0], addressID, "AddressID is different from expected")
	assert.Nil(t, err, "Error occurred when Address was instantiated")
}
