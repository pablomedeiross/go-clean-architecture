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

	if userTest.GetName() != nameExpected {
		t.Errorf("Invalid name returned from GetName")
	}
}

func TestAddAddressInUserWithoutAddress(t *testing.T) {

	address, _ := domain.NewAddress("test", 2, "test", 12345567)
	userTest, _ := domain.NewUser("name", "email@gmail.com", 10, nil)

	userTest.AddAddress(address)
}

//TODO:
func TestAddAddressInUserWithSameAddress(t *testing.T) {

	address, _ := domain.NewAddress("test", 2, "test", 12345567)
	user, err := domain.NewUser("name", "email@gmail.com", 10, append([]domain.Address{}, address))

	user.AddAddress(address)

	if err == nil {
		t.Errorf("Error")
	}
}
