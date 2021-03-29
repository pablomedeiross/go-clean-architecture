package db_test

import (
	"testing"
	"user-api/adapter/db"
	"user-api/adapter/db/double"
	"user-api/entity/user"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

var expectedUser user.User
var inputUser user.User

func init() {

	inputUser, _ = user.
		NewBuilder().
		Name("name").
		Email("email").
		Age(12).
		Build()

	expectedUser, _ = user.
		NewBuilder().
		Name(inputUser.Name()).
		Email(inputUser.Email()).
		Age(inputUser.Age()).
		Id("id").
		Build()
}

func TestNewUserRepository(t *testing.T) {

	dbGatewayDouble := double.NewDBGateway(nil, nil)

	repository, err := db.NewUserRepository(dbGatewayDouble)

	assert.Nil(t, err)
	assert.NotNil(t, repository)
}

func TestNewUserRepositoryReturnErrorWithoutDBGateway(t *testing.T) {

	repository, err := db.NewUserRepository(nil)

	assert.Nil(t, repository)
	assert.Error(t, err)
}

func TestSaveUser(t *testing.T) {

	dbGatewayDouble := double.NewDBGateway(

		func(user db.User) (db.User, error) {

			return db.User{
					expectedUser.Id(),
					expectedUser.Name(),
					expectedUser.Email(),
					expectedUser.Age(),
					expectedUser.AddressesIds(),
				},

				nil
		},

		nil,
	)

	repo, _ := db.NewUserRepository(dbGatewayDouble)
	user, err := repo.Save(inputUser)

	assert.Nil(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestSaveUserErrorInDbGateway(t *testing.T) {

	dbGatewayDouble := double.NewDBGateway(

		func(user db.User) (db.User, error) {
			return db.User{}, errors.New("Error")
		},
		nil,
	)

	repo, _ := db.NewUserRepository(dbGatewayDouble)
	user, err := repo.Save(inputUser)

	assert.Nil(t, user)
	assert.Error(t, err)
}

func TestSaveUserErrorToTheCreateUser(t *testing.T) {

	dbGatewayDouble := double.NewDBGateway(

		func(user db.User) (db.User, error) {
			return db.User{
					Id:           "id",
					Name:         "",
					Email:        "email",
					Age:          12,
					AddressesIds: []string{}},

				nil
		},
		nil,
	)

	repo, _ := db.NewUserRepository(dbGatewayDouble)
	user, err := repo.Save(inputUser)

	assert.Nil(t, user)
	assert.Error(t, err)
}
