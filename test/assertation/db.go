package assertation

import (
	"testing"
	"user-api/db/inmemory"
	"user-api/test/helper"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// To use this function is necessary that a instance of mongoDB be in execution
func AssertThatUserExistsInDB(t *testing.T, id string, dbHelper *inmemory.InMemoryMongoDB) {

	primitiveId, _ := primitive.ObjectIDFromHex(id)
	userReturned, err := helper.FindUserById(dbHelper, primitiveId)

	assert.Equal(t, userReturned.Id.Hex(), id, "User don't exists in DB, id: "+id)
	assert.NoError(t, err)
}

// To use this function is necessary that a instance of mongoDB be in execution
func AssertThatUserDontExistsInDB(t *testing.T, name string, dbHelper *inmemory.InMemoryMongoDB) {

	userReturned, err := helper.FindUserByName(dbHelper, name)

	assert.Empty(t, userReturned)
	assert.Error(t, err)
}
