package db_test

import (
	"context"
	"testing"
	adapter "user-api/adapter/db"
	"user-api/external/db"
	"user-api/external/db/test/assertation"
	"user-api/external/db/test/helper"

	"github.com/stretchr/testify/assert"
)

var dbHelper *helper.DBHelper
var userTest adapter.User

func init() {

	dbHelper = helper.NewMongoDBHelper()

	userTest = adapter.User{
		Name:  "test",
		Email: "email",
		Age:   13,
	}

}

func TestNewDBGateway(t *testing.T) {

	dbHelper.StartMongoDB()
	defer dbHelper.StopMongoDB()

	dbGateway, err := db.NewNoSQLDB(dbHelper.DatabaseURI(), dbHelper.DatabaseName())
	assert.NotNil(t, dbGateway)
	assert.Nil(t, err)
}

func TestNewDBGatewayError(t *testing.T) {

	dbHelper.StartMongoDB()
	defer dbHelper.StopMongoDB()

	dbGateway, err := db.NewNoSQLDB(dbHelper.DatabaseURI(), dbHelper.DatabaseName())
	assert.NotNil(t, dbGateway)
	assert.Nil(t, err)
}

func TestSaveUser(t *testing.T) {

	dbHelper.StartMongoDB()
	defer dbHelper.StopMongoDB()

	dbGateway, err := db.NewNoSQLDB(dbHelper.DatabaseURI(), dbHelper.DatabaseName())
	id, err := dbGateway.SaveUser(context.Background(), userTest)

	assertation.AssertThatUserExistsInDB(t, id, dbHelper)
	assert.Nil(t, err)
}

func TestFindUserByName(t *testing.T) {

	err := dbHelper.StartMongoDB()
	defer dbHelper.StopMongoDB()

	err = dbHelper.InsertUser(userTest)

	dbGateway, err := db.NewNoSQLDB(dbHelper.DatabaseURI(), dbHelper.DatabaseName())
	usr, err := dbGateway.FindUserByName(context.Background(), userTest.Name)

	assert.Nil(t, err)
	assertation.AssertThatUserExistsInDB(t, usr.Id, dbHelper)
	assertation.AssertThatUserEqualWithouId(t, userTest, usr)
}

func TestFindUserByNameWithoutNoneExistentName(t *testing.T) {

	err := dbHelper.StartMongoDB()
	defer dbHelper.StopMongoDB()

	err = dbHelper.InsertUser(userTest)

	dbGateway, err := db.NewNoSQLDB(dbHelper.DatabaseURI(), dbHelper.DatabaseName())
	usr, err := dbGateway.FindUserByName(context.Background(), "nonexistent name")

	assert.Error(t, err)
	assert.Empty(t, usr)
}
