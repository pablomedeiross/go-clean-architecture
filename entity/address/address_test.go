package address_test

import (
	"strconv"
	"testing"
	"user-api/entity/address"

	"github.com/stretchr/testify/assert"
)

const idTestAddress string = "ID"
const streetTestAddress string = "Cool Street"
const numberTestAddress int = 12356
const neighborhoodTestAddress string = "Cool Neighborhood"
const zipcodeTestAddress int = 31298031

const msgErrorInvalidParamAddress string = "Param different from expected when build Address"

var addressTest address.Address

func init() {
	addressTest, _ = address.
		NewBuilder().
		Id(idTestAddress).
		Street(streetTestAddress).
		Number(numberTestAddress).
		Neighborhood(neighborhoodTestAddress).
		Zipcode(zipcodeTestAddress).
		Build()
}

func TestGetID(t *testing.T) {

	idExpected := idTestAddress
	assert.Equal(t, addressTest.Id(), idExpected, "Error GetId() return another value")
}

func TestGetStreet(t *testing.T) {

	streetExpected := streetTestAddress

	msgError := "Invalid street in GetStreet() test. Actual : " +
		addressTest.Street() +
		"expected: " +
		streetExpected

	assert.Equal(t, streetExpected, addressTest.Street(), msgError)
}

func TestGetNumber(t *testing.T) {

	numberExpected := numberTestAddress

	msgError := "Invalid number in GetNumber() test. Actual : " +
		strconv.Itoa(addressTest.Number()) +
		"expected: " +
		strconv.Itoa(numberExpected)

	assert.Equal(t, numberExpected, addressTest.Number(), msgError)
}

func TestGetNeighborhood(t *testing.T) {

	neighborhoodExpected := neighborhoodTestAddress

	msgError :=
		"Invalid neighborhood in GetNeighborhood() test. Actual : " +
			addressTest.Neighborhood() +
			"expected: " +
			neighborhoodExpected

	assert.Equal(t, neighborhoodExpected, addressTest.Neighborhood(), msgError)
}

func TestGetZipCode(t *testing.T) {

	zipCodeExpected := zipcodeTestAddress

	msgError := "Invalid number in GetZipCode() test. Actual : " +
		strconv.Itoa(addressTest.Zipcode()) +
		"expected: " +
		strconv.Itoa(zipCodeExpected)

	assert.Equal(t, zipCodeExpected, addressTest.Zipcode(), msgError)
}
