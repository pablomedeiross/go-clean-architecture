package assertation

import (
	"testing"
	"user-api/adapter/db"
	"user-api/entity/user"

	"github.com/stretchr/testify/assert"
)

func UserEntityEqualDB(t *testing.T, entity user.User, db db.User) {

	assert.Equal(t, entity.Id(), db.Id.Hex())
	UserEntityEqualDBWithoutId(t, entity, db)
}

func UsersEqualWithoutId(t *testing.T, expected user.User, actual user.User) {

	assert.Equal(t, expected.Name(), actual.Name())
	assert.Equal(t, expected.Email(), actual.Email())
	assert.Equal(t, expected.Age(), actual.Age())
	assert.Equal(t, expected.AddressesIds(), actual.AddressesIds())
}

func UserEntityEqualDBWithoutId(t *testing.T, entity user.User, db db.User) {

	assert.Equal(t, entity.Name(), db.Name)
	assert.Equal(t, entity.Email(), db.Email)
	assert.Equal(t, entity.Age(), db.Age)
	assert.Equal(t, entity.AddressesIds(), db.AddressesIds)
}
