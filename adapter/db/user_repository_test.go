package db_test

import (
	"context"
	"testing"
	"user-api/adapter/db"
	"user-api/adapter/db/test/assertation"
	"user-api/adapter/db/test/double"
	"user-api/entity/user"
	"user-api/usecase"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

	dbDouble := double.NewNoSQLDB(nil, nil)

	repository, err := db.NewUserRepository(&dbDouble)

	assert.NoError(t, err)
	assert.NotNil(t, repository)
}

func TestNewUserRepositoryReturnErrorWithoutDBGateway(t *testing.T) {

	repository, err := db.NewUserRepository(nil)

	assert.Nil(t, repository)
	assert.Error(t, err)
}

func TestSaveUser(t *testing.T) {

	dbDouble := double.NewNoSQLDB(

		func(ctx context.Context, user db.User) (primitive.ObjectID, error) {
			return primitive.NewObjectID(), nil
		},

		nil,
	)

	repo, _ := db.NewUserRepository(&dbDouble)
	returnedUser, err := repo.Save(context.Background(), inputUser)

	assert.NoError(t, err)
	assert.NotEmpty(t, expectedUser.Id())
	assertation.UsersEqualWithoutId(t, inputUser, returnedUser)
}

func TestSaveUserErrorInDbGateway(t *testing.T) {

	dbDouble := double.NewNoSQLDB(

		func(ctx context.Context, user db.User) (primitive.ObjectID, error) {
			return primitive.NilObjectID, errors.New("Error")
		},
		nil,
	)

	repo, _ := db.NewUserRepository(&dbDouble)
	returnedUser, err := repo.Save(context.Background(), inputUser)

	assert.Nil(t, returnedUser)
	assert.Error(t, err)
}

func TestFindUserByNameReturnUser(t *testing.T) {

	primitiveId, _ := primitive.ObjectIDFromHex(expectedUser.Id())

	expectedUserDB := db.User{
		Id:    primitiveId,
		Name:  inputUser.Name(),
		Email: inputUser.Email(),
		Age:   inputUser.Age(),
	}

	dbDouble := double.NewNoSQLDB(

		nil,
		func(ctx context.Context, name string) (db.User, error) {
			return expectedUserDB, nil
		},
	)

	repo, _ := db.NewUserRepository(&dbDouble)
	returnedUser, err := repo.FindByName(context.Background(), inputUser.Name())

	assert.NoError(t, err)
	assertation.UserEntityEqualDB(t, returnedUser, expectedUserDB)
}

func TestFindUserByNameReturnError(t *testing.T) {

	dbDouble := double.NewNoSQLDB(

		nil,
		func(ctx context.Context, name string) (db.User, error) {
			return db.User{}, mongo.ErrNoDocuments
		},
	)

	repo, _ := db.NewUserRepository(&dbDouble)
	returnedUser, err := repo.FindByName(context.Background(), inputUser.Name())

	assert.IsType(t, err, usecase.NewUserDontExistError(inputUser.Name()))
	assert.Empty(t, returnedUser)
}
