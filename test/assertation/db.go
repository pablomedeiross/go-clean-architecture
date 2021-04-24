package assertation

import (
	"testing"
	"user-api/db/inmemory"
	"user-api/test/helper"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// To use this function is necessary that a instance of mongoDB be in execution
func AssertThatUserExistsInDB(t *testing.T, id primitive.ObjectID, dbHelper *inmemory.InMemoryMongoDB) {

	userReturned, err := helper.FindUserById(dbHelper, id)

	assert.Equal(t, userReturned.Id, id, "User don't exists in DB, id: "+id.Hex())
	assert.NoError(t, err)
}

// To use this function is necessary that a instance of mongoDB be in execution
func AssertThatUserDontExistsInDB(t *testing.T, name string, dbHelper *inmemory.InMemoryMongoDB) {

	userReturned, err := helper.FindUserByName(dbHelper, name)

	assert.Empty(t, userReturned)
	assert.Error(t, err)
}
