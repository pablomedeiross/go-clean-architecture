package address_test

import (
	"testing"
	"user-api/entity/address"

	"github.com/stretchr/testify/assert"
)

func TestInvalidBuildAddress(t *testing.T) {

	var invalidBuilders []address.Builder = []address.Builder{
		address.NewBuilder().Id(idTest).Street(streetTest).Number(numberTest).Neighborhood(neighborhoodTest),
		address.NewBuilder().Id(idTest).Street(streetTest).Number(numberTest).Zipcode(zipcodeTest),
		address.NewBuilder().Id(idTest).Street(streetTest).Neighborhood(neighborhoodTest).Zipcode(zipcodeTest),
		address.NewBuilder().Id(idTest).Number(numberTest).Neighborhood(neighborhoodTest).Zipcode(zipcodeTest),
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
		Id(idTest).
		Street(streetTest).
		Number(numberTest).
		Neighborhood(neighborhoodTest).
		Zipcode(zipcodeTest).
		Build()

	validateRequestParams(t, address, err)
	assert.Equal(t, idTest, address.Id(), msgErrorInvalidParam)
}

func TestBuildAddressWithoutID(t *testing.T) {

	address, err := address.
		NewBuilder().
		Street(streetTest).
		Number(numberTest).
		Neighborhood(neighborhoodTest).
		Zipcode(zipcodeTest).
		Build()

	validateRequestParams(t, address, err)
}

func validateRequestParams(t *testing.T, address address.Address, err error) {

	assert.Nil(t, err, "Error non nil when create Address from Builder")
	assert.Equal(t, streetTest, address.Street(), msgErrorInvalidParam)
	assert.Equal(t, numberTest, address.Number(), msgErrorInvalidParam)
	assert.Equal(t, neighborhoodTest, address.Neighborhood(), msgErrorInvalidParam)
	assert.Equal(t, zipcodeTest, address.Zipcode(), msgErrorInvalidParam)
}
