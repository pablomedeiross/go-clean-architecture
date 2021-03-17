package address_test

import (
	"testing"
	"user-api/entity/address"

	"github.com/stretchr/testify/assert"
)

var msgErrorInvalidParam string = "Param different from expected when build Address"

func TestInvalidBuildAddress(t *testing.T) {

	var invalidBuilders []address.Builder = []address.Builder{
		address.NewBuilder().ID(Id).Street(Street).Number(Number).Neighborhood(Neighborhood),
		address.NewBuilder().ID(Id).Street(Street).Number(Number).ZipCode(Zipcode),
		address.NewBuilder().ID(Id).Street(Street).Neighborhood(Neighborhood).ZipCode(Zipcode),
		address.NewBuilder().ID(Id).Number(Number).Neighborhood(Neighborhood).ZipCode(Zipcode),
	}

	for _, builder := range invalidBuilders {

		address, err := builder.Build()
		assert.Nil(t, address, "Address non nil when build without a requested param")
		assert.Error(t, err, "Error is nil when build address without a requested param")
	}
}

func TestBuildAddressWithAllParams(t *testing.T) {
	address, err := address.NewBuilder().ID(Id).Street(Street).Number(Number).Neighborhood(Neighborhood).ZipCode(Zipcode).Build()
	validRequestParams(t, address, err)
	assert.EqualValues(t, Id, address.GetID(), msgErrorInvalidParam)
}

func TestBuildAddressWithoutID(t *testing.T) {

	address, err := address.NewBuilder().Street(Street).Number(Number).Neighborhood(Neighborhood).ZipCode(Zipcode).Build()
	validRequestParams(t, address, err)
}

func validRequestParams(t *testing.T, address address.Address, err error) {

	assert.Nil(t, err, "Error non nil when create Address from Builder")
	assert.EqualValues(t, Street, address.Street(), msgErrorInvalidParam)
	assert.EqualValues(t, Number, address.Number(), msgErrorInvalidParam)
	assert.EqualValues(t, Neighborhood, address.Neighborhood(), msgErrorInvalidParam)
	assert.EqualValues(t, Zipcode, address.ZipCode(), msgErrorInvalidParam)
}
