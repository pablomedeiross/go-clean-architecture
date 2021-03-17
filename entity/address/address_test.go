package address_test

import (
	"strconv"
	"testing"
	"user-api/entity/address"

	"github.com/stretchr/testify/assert"
)

var addressTest address.Address

func init() {
	addressTest, _ = address.NewBuilder().ID(Id).Street(Street).Number(Number).Neighborhood(Neighborhood).ZipCode(Zipcode).Build()
}

func TestGetID(t *testing.T) {

	idExpected := Id
	assert.EqualValues(t, addressTest.GetID(), idExpected, "Error GetId() return another value")
}

func TestGetStreet(t *testing.T) {

	streetExpected := Street

	msgError := "Invalid street in GetStreet() test. Actual : " +
		addressTest.Street() +
		"expected: " +
		streetExpected

	assert.EqualValues(t, streetExpected, addressTest.Street(), msgError)
}

func TestGetNumber(t *testing.T) {

	numberExpected := Number

	msgError := "Invalid number in GetNumber() test. Actual : " +
		strconv.Itoa(addressTest.Number()) +
		"expected: " +
		strconv.Itoa(numberExpected)

	assert.EqualValues(t, numberExpected, addressTest.Number(), msgError)
}

func TestGetNeighborhood(t *testing.T) {

	neighborhoodExpected := Neighborhood

	msgError :=
		"Invalid neighborhood in GetNeighborhood() test. Actual : " +
			addressTest.Neighborhood() +
			"expected: " +
			neighborhoodExpected

	assert.EqualValues(t, neighborhoodExpected, addressTest.Neighborhood(), msgError)
}

func TestGetZipCode(t *testing.T) {

	zipCodeExpected := Zipcode

	msgError := "Invalid number in GetZipCode() test. Actual : " +
		strconv.Itoa(addressTest.ZipCode()) +
		"expected: " +
		strconv.Itoa(zipCodeExpected)

	assert.EqualValues(t, zipCodeExpected, addressTest.ZipCode(), msgError)
}
