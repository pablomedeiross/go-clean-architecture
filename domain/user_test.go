package domain_test

import (
	"domain/domain"
	"errors"
	"reflect"
	"strconv"
	"testing"
)

func TestNewUserValidation(t *testing.T) {

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

		_, errorActual := domain.NewUser(names[i], emails[i], ages[i], nil)

		if !reflect.DeepEqual(errorActual, errorExpected) {
			t.Errorf("Invalid error message !" + errorActual.Error())
		}
	}
}

func TestGetName(t *testing.T) {

	nameExpected := "NameForTest"
	userTest, _ := domain.NewUser(nameExpected, "email@gmail.com", 10, nil)

	if userTest.Name() != nameExpected {
		t.Errorf("Invalid name returned from Name()")
	}
}

func TestGetEmail(t *testing.T) {

	emailExpected := "email@gmail.com"
	userTest, _ := domain.NewUser("name", emailExpected, 10, nil)

	if userTest.Email() != emailExpected {
		t.Errorf("Invalid email returned from Email()")
	}
}

func TestGetAge(t *testing.T) {

	ageExpected := 20
	userTest, _ := domain.NewUser("name", "email@gmail.com", ageExpected, nil)

	if userTest.Age() != ageExpected {
		t.Errorf("Invalid age returned from Age()")
	}
}

func TestGetAdresses(t *testing.T) {

	address, _ := domain.NewAddress("test", 2, "test", 12345567)
	addressesExpected := append([]domain.Address{}, address)

	userTest, _ := domain.NewUser("name", "email@gmail.com", 12, addressesExpected)

	if !reflect.DeepEqual(userTest.Addresses(), addressesExpected) {
		t.Errorf("Invalid addresses returned from Addresses()")
	}
}

func TestAddAddressInUserWithoutAddress(t *testing.T) {

	address, _ := domain.NewAddress("test", 2, "test", 12345567)
	userTest, _ := domain.NewUser("name", "email@gmail.com", 10, nil)

	err := userTest.AddAddress(address)

	if err != nil {
		t.Errorf("Error error returning when addAddress in user: " + err.Error())
	}
}

//TODO:
func TestAddAddressInUserWithSameAddress(t *testing.T) {

	address, _ := domain.NewAddress("test", 2, "test", 12345567)
	user, _ := domain.NewUser("name", "email@gmail.com", 10, append([]domain.Address{}, address))

	err := user.AddAddress(address)

	if err == nil {
		t.Errorf("Error AddAddress add same address in User")
	}
}
