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

// TODO: Test
// func TestAddAddressInUserWithoutAddress(t *testing.T) {

// 	address := address
// 	userTest, _ := domain.NewUser(nameExpected, "email@gmail.com", 10, nil)

// }
