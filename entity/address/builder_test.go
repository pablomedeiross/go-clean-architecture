package address_test

import (
	"testing"
	"user-api/entity/address"

	"github.com/stretchr/testify/assert"
)

const idBuilderTest string = "ID"
const streetBuilderTest string = "Cool Street"
const numberBuilderTest int = 12356
const neighborhoodBuilderTest string = "Cool Neighborhood"
const zipcodeBuilderTest int = 31298031

const msgErrorInvalidParamBuild string = "Param different from expected when build Address"

func TestInvalidBuildAddress(t *testing.T) {

	var invalidBuilders []address.Builder = []address.Builder{
		address.NewBuilder().Id(idBuilderTest).Street(streetBuilderTest).Number(numberBuilderTest).Neighborhood(neighborhoodBuilderTest),
		address.NewBuilder().Id(idBuilderTest).Street(streetBuilderTest).Number(numberBuilderTest).Zipcode(zipcodeBuilderTest),
		address.NewBuilder().Id(idBuilderTest).Street(streetBuilderTest).Neighborhood(neighborhoodBuilderTest).Zipcode(zipcodeBuilderTest),
		address.NewBuilder().Id(idBuilderTest).Number(numberBuilderTest).Neighborhood(neighborhoodBuilderTest).Zipcode(zipcodeBuilderTest),
	}

	for _, builder := range invalidBuilders {

		address, err := builder.Build()
		assert.Nil(t, address, "Address non nil when build without a requested param")
		assert.Error(t, err, "Error is nil when build address without a requested param")
	}
}

func TestBuildAddressWithAllParams(t *testing.T) {

	address, err := address.
		NewBuilder().
		Id(idBuilderTest).
		Street(streetBuilderTest).
		Number(numberBuilderTest).
		Neighborhood(neighborhoodBuilderTest).
		Zipcode(zipcodeBuilderTest).
		Build()

	validateRequestParams(t, address, err)
	assert.Equal(t, idBuilderTest, address.Id(), msgErrorInvalidParamBuild)
}

func TestBuildAddressWithoutID(t *testing.T) {

	address, err := address.
		NewBuilder().
		Street(streetBuilderTest).
		Number(numberBuilderTest).
		Neighborhood(neighborhoodBuilderTest).
		Zipcode(zipcodeBuilderTest).
		Build()

	validateRequestParams(t, address, err)
}

func validateRequestParams(t *testing.T, address address.Address, err error) {

	assert.Nil(t, err, "Error non nil when create Address from Builder")
	assert.Equal(t, streetBuilderTest, address.Street(), msgErrorInvalidParamBuild)
	assert.Equal(t, numberBuilderTest, address.Number(), msgErrorInvalidParamBuild)
	assert.Equal(t, neighborhoodBuilderTest, address.Neighborhood(), msgErrorInvalidParamBuild)
	assert.Equal(t, zipcodeBuilderTest, address.Zipcode(), msgErrorInvalidParamBuild)
}
