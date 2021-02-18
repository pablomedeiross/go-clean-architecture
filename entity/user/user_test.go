package user_test

import (
	"testing"
	"errors"
	"strconv"
	"user-api/entity/user"
	"github.com/stretchr/testify/assert"
)

func TestValidationNew(t *testing.T) {

	names, emails, ages :=
		[]string{"", "name2", "name3"},
		[]string{"test@hotmail.com", "", "test4@live.com"},
		[]int{10, 11, 0}

	for i, _ := range names {

		errorExpected := errors.New(
			"Error creating new User with arguments : " +
				names[i] + ", " +
				emails[i] + ", " +
				strconv.Itoa(ages[i]))

		_, errorActual := user.New(names[i], emails[i], ages[i])
		assert.Equal(t, errorExpected, errorActual, "Invalid error message !" + errorActual.Error())
	}
}

func TestGetName(t *testing.T) {

	nameExpected := "NameForTest"
	userTest, _ := user.New(nameExpected, "email@gmail.com", 10)
	assert.Equal(t, nameExpected, userTest.Name(), "Invalid name returned from Name()")
}

func TestGetEmail(t *testing.T) {

	emailExpected := "email@gmail.com"
	userTest, _ := user.New("name", emailExpected, 10)
	assert.Equal(t, emailExpected, userTest.Email(), "Invalid email returned from Email()")
}

func TestGetAge(t *testing.T) {

	ageExpected := 20
	userTest, _ := user.New("name", "email@gmail.com", ageExpected)
	assert.Equal(t, ageExpected, userTest.Age(), "Invalid age returned from Age()")
}

func TestGetAdressesIDs(t *testing.T) {

	addressesIDsExpected := []string{"122445"}
	userTest, _ := user.NewPersisted(1, "name", "email@gmail.com", 12, addressesIDsExpected)
	assert.Equal(t, addressesIDsExpected, userTest.AddressesIDs(), "Invalid addresses returned from AddressesIDs()")
}

func TestAddAddressIDInUserWithoutAddressId(t *testing.T) {

	addressID := "56789"
	userTest, _ := user.New("name", "email@gmail.com", 10)

	err := userTest.AddAddressID(addressID)
	assert.Nil(t, err)
}

//TODO:
func TestAddAddressInUserWithSameAddress(t *testing.T) {

	addressId := "1223"
	user, _ := user.NewPersisted(1, "name", "email@gmail.com", 10, []string{addressId})
	err := user.AddAddressID(addressId)

	msgErrorExpected := "AddressId already exists in User: " + "name"
	assert.EqualError(t, err, msgErrorExpected, "Error AddAddress add same address in User")
}
