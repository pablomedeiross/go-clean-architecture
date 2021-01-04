package domain_test

import (
	"domain/domain"
	"errors"
	"reflect"
	"strconv"
	"testing"
)

func TestNewAddressValidations(t *testing.T) {

	streets, numbers, neighborhoods, zipCodes :=
		[]string{"", "street2", "street3", "street4"},
		[]int{12345, 0, 5313, 14664},
		[]string{"neighborhood1", "neighborhood2", "", "neighborhood4"},
		[]int{10450103, 10455112, 90750103, 0}

	for i, _ := range streets {

		errorExpected := errors.New(
			"Error creating new Address with arguments : " +
				streets[i] + ", " +
				strconv.Itoa(numbers[i]) + ", " +
				neighborhoods[i] + ", " +
				strconv.Itoa(zipCodes[i]))

		addressActual, errorActual := domain.NewAddress(streets[i], numbers[i], neighborhoods[i], zipCodes[i])

		if addressActual != nil {
			t.Errorf("Invalid address instatiation: " + addressActual.GetStreet())
		}

		if errorActual == nil || !reflect.DeepEqual(errorActual, errorExpected) {
			t.Errorf("Invalid error message ! ---" + errorActual.Error())
		}
	}
}

func TestGetStreet(t *testing.T) {

	streetExpected := "test street"

	addressActual, _ := domain.NewAddress(streetExpected, 1112, "neighborhood", 123452341)

	if addressActual.GetStreet() != streetExpected {
		t.Errorf(
			"Invalid street in GetStreet() test. Actual : " +
				addressActual.GetStreet() +
				"expected: " +
				streetExpected)
	}
}

func TestGetNumber(t *testing.T) {

	numberExpected := 4312

	addressActual, _ := domain.NewAddress("test", numberExpected, "neighborhood", 123452341)

	if addressActual.GetNumber() != numberExpected {
		t.Errorf(
			"Invalid number in GetNumber() test. Actual : " +
				strconv.Itoa(addressActual.GetNumber()) +
				"expected: " +
				strconv.Itoa(numberExpected))
	}
}

func TestGetNeighborhood(t *testing.T) {

	neighborhoodExpected := "test neighborhood"

	addressActual, _ := domain.NewAddress("test", 1112, neighborhoodExpected, 123452341)

	if addressActual.GetNeighborhood() != neighborhoodExpected {
		t.Errorf(
			"Invalid neighborhood in GetNeighborhood() test. Actual : " +
				addressActual.GetNeighborhood() +
				"expected: " +
				neighborhoodExpected)
	}
}

func TestGetZipCode(t *testing.T) {

	zipCodeExpected := 4312

	addressActual, _ := domain.NewAddress("test", 124, "neighborhood", zipCodeExpected)

	if addressActual.GetZipCode() != zipCodeExpected {
		t.Errorf(
			"Invalid number in GetZipCode() test. Actual : " +
				strconv.Itoa(addressActual.GetZipCode()) +
				"expected: " +
				strconv.Itoa(zipCodeExpected))
	}
}
