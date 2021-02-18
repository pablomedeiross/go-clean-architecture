package address_test

import (
	"testing"
	"user-api/entity/address"

	"github.com/stretchr/testify/assert"
)

func TestNewAddressValidations(t *testing.T) {

	streets, numbers, neighborhoods, zipCodes :=
		[]string{"", "street2", "street3", "street4"},
		[]int{12345, 0, 5313, 14664},
		[]string{"neighborhood1", "neighborhood2", "", "neighborhood4"},
		[]int{10450103, 10455112, 90750103, 0}

	for i, _ := range streets {

		// errorExpected :=
		// 	"Error creating new Address with arguments : " +
		// 		streets[i] + ", " +
		// 		strconv.Itoa(numbers[i]) + ", " +
		// 		neighborhoods[i] + ", " +
		// 		strconv.Itoa(zipCodes[i])

		addressActual, errorActual := address.New(streets[i], numbers[i], neighborhoods[i], zipCodes[i])
		// assert.Nil(t, addressActual, "Address created without requested parameters")
		assert.EqualError(t, errorActual, "teste", "Address created without requested")
	}
}

// func TestGetStreet(t *testing.T) {

// 	streetExpected := "test street"

// 	addressActual, _ := address.New(streetExpected, 1112, "neighborhood", 123452341)

// 	if addressActual.Street() != streetExpected {
// 		t.Errorf(
// 			"Invalid street in GetStreet() test. Actual : " +
// 				addressActual.Street() +
// 				"expected: " +
// 				streetExpected)
// 	}
// }

// func TestGetNumber(t *testing.T) {

// 	numberExpected := 4312

// 	addressActual, _ := address.New("test", numberExpected, "neighborhood", 123452341)

// 	if addressActual.Number() != numberExpected {
// 		t.Errorf(
// 			"Invalid number in GetNumber() test. Actual : " +
// 				strconv.Itoa(addressActual.Number()) +
// 				"expected: " +
// 				strconv.Itoa(numberExpected))
// 	}
// }

// func TestGetNeighborhood(t *testing.T) {

// 	neighborhoodExpected := "test neighborhood"

// 	addressActual, _ := address.New("test", 1112, neighborhoodExpected, 123452341)

// 	if addressActual.Neighborhood() != neighborhoodExpected {
// 		t.Errorf(
// 			"Invalid neighborhood in GetNeighborhood() test. Actual : " +
// 				addressActual.Neighborhood() +
// 				"expected: " +
// 				neighborhoodExpected)
// 	}
// }

// func TestGetZipCode(t *testing.T) {

// 	zipCodeExpected := 4312

// 	addressActual, _ := address.New("test", 124, "neighborhood", zipCodeExpected)

// 	if addressActual.ZipCode() != zipCodeExpected {
// 		t.Errorf(
// 			"Invalid number in GetZipCode() test. Actual : " +
// 				strconv.Itoa(addressActual.ZipCode()) +
// 				"expected: " +
// 				strconv.Itoa(zipCodeExpected))
// 	}
// }
