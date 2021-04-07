package assertation

import (
	"testing"
	"user-api/adapter/db"
	"user-api/external/db/memory"
	test_helper "user-api/external/db/test/helper"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AssertThatUserEqualWithouId(t *testing.T, expected db.User, actual db.User) {

	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.Email, actual.Email)
	assert.Equal(t, expected.Age, actual.Age)
	assert.Equal(t, expected.AddressesIds, actual.AddressesIds)
}

// To use this function is necessary that a instance of mongoDB be in execution
func AssertThatUserExistsInDB(t *testing.T, id primitive.ObjectID, dbHelper *memory.InMemoryMongoDB) {

	userReturned, err := test_helper.FindUserById(dbHelper, id)

	assert.Equal(t, userReturned.Id, id, "User don't exists in DB, id: "+id.Hex())
	assert.NoError(t, err)
}
