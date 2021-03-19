package address_test

import (
	"strconv"
	"testing"
	"user-api/entity/address"

	"github.com/stretchr/testify/assert"
)

var addressTest address.Address

func init() {
	addressTest, _ = address.
		NewBuilder().
		Id(idTest).
		Street(streetTest).
		Number(numberTest).
		Neighborhood(neighborhoodTest).
		Zipcode(zipcodeTest).
		Build()
}

func TestGetID(t *testing.T) {

	idExpected := idTest
	assert.Equal(t, addressTest.Id(), idExpected, "Error GetId() return another value")
}

func TestGetStreet(t *testing.T) {

	streetExpected := streetTest

	msgError := "Invalid street in GetStreet() test. Actual : " +
		addressTest.Street() +
		"expected: " +
		streetExpected

	assert.Equal(t, streetExpected, addressTest.Street(), msgError)
}

func TestGetNumber(t *testing.T) {

	numberExpected := numberTest

	msgError := "Invalid number in GetNumber() test. Actual : " +
		strconv.Itoa(addressTest.Number()) +
		"expected: " +
		strconv.Itoa(numberExpected)

	assert.Equal(t, numberExpected, addressTest.Number(), msgError)
}

func TestGetNeighborhood(t *testing.T) {

	neighborhoodExpected := neighborhoodTest

	msgError :=
		"Invalid neighborhood in GetNeighborhood() test. Actual : " +
			addressTest.Neighborhood() +
			"expected: " +
			neighborhoodExpected

	assert.Equal(t, neighborhoodExpected, addressTest.Neighborhood(), msgError)
}

func TestGetZipCode(t *testing.T) {

	zipCodeExpected := zipcodeTest

	msgError := "Invalid number in GetZipCode() test. Actual : " +
		strconv.Itoa(addressTest.Zipcode()) +
		"expected: " +
		strconv.Itoa(zipCodeExpected)

	assert.Equal(t, zipCodeExpected, addressTest.Zipcode(), msgError)
}
