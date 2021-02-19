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

		userTest, errorActual := user.New(names[i], emails[i], ages[i])
		assert.Nil(t, userTest, "Error, user created withou requested parameters!")
		assert.Equal(t, errorExpected, errorActual, "Invalid error message !" + errorActual.Error())
	}
}

func TestValidationNewPersisted(t *testing.T) {

	ids, names, emails, ages, addressesIds :=
		[]int{1234, 12345678910, 12345678910, 12345678910},
		[]string{"name2", "", "name3", "name4"},
		[]string{"test@gmail.com","test@hotmail.com", "", "test4@live.com","test4@bol.com"},
		[]int{10, 11, 192, 0, 181},
		[][]string{{"testAddress1"},{"testAddress2"},{"testAddress3"},{"testAddress4"}, nil}

	for i, _ := range names {

		errorExpected := errors.New(
			"Error creating new User with arguments : " +
			names[i] + ", " +
			emails[i] + ", " +
			strconv.Itoa(ages[i]))

		userTest, errorActual := user.NewPersisted(ids[i], names[i], emails[i], ages[i], addressesIds[i])
		assert.Nil(t, userTest, "Error, user created withou requested parameters!")
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
	userTest, _ := user.NewPersisted(12345678910, "name", "email@gmail.com", 12, addressesIDsExpected)
	assert.Equal(t, addressesIDsExpected, userTest.AddressesIDs(), "Invalid addresses returned from AddressesIDs()")
}

func TestAddAddressIDInUserWithoutAddressId(t *testing.T) {

	addressID := "56789"
	userTest, _ := user.New("name", "email@gmail.com", 10)

	err := userTest.AddAddressID(addressID)
	assert.Nil(t, err)
}

func TestAddAddressInUserWithSameAddress(t *testing.T) {

	addressId := "1223"
	user, _ := user.NewPersisted(12345678910, "name", "email@gmail.com", 10, []string{addressId})
	err := user.AddAddressID(addressId)

	msgErrorExpected := "AddressId already exists in User: " + "name"
	assert.EqualError(t, err, msgErrorExpected, "Error AddAddress add same address in User")
}
