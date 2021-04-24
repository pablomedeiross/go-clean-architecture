package db_test

import (
	"context"
	"testing"
	adapter "user-api/adapter/db"
	memory "user-api/db/inmemory"
	"user-api/external/db"
	"user-api/test/assertation"

	"user-api/test/helper"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DBSuite struct {
	suite.Suite
	mongoDB     memory.InMemoryMongoDB
	userForTest adapter.User
}

func (suite *DBSuite) SetupSuite() {

	suite.mongoDB = *memory.NewInMemoryMongoDB()
	suite.mongoDB.Start()

	suite.userForTest = adapter.User{
		Name:  "test",
		Email: "email",
		Age:   13,
	}
}

func (suite *DBSuite) TearDownSuite() {
	suite.mongoDB.Stop()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(DBSuite))
}

func (suite *DBSuite) TestNewDBGateway() {

	dbGateway, err := db.NewNoSQLDB(suite.mongoDB.URI(), suite.mongoDB.Name())
	suite.NotNil(dbGateway)
	suite.Nil(err)
}

func (suite *DBSuite) TestNewDBGatewayError() {

	dbGateway, err := db.NewNoSQLDB("", "")
	suite.Error(err)
	suite.Nil(dbGateway)
}

func (suite *DBSuite) TestSaveUser() {

	dbGateway, err := db.NewNoSQLDB(suite.mongoDB.URI(), suite.mongoDB.Name())
	assert.NoError(suite.T(), err)

	id, err := dbGateway.SaveUser(context.Background(), suite.userForTest)

	assertation.AssertThatUserExistsInDB(suite.T(), id.Hex(), &suite.mongoDB)
	suite.Nil(err)
}

func (suite *DBSuite) TestFindUserByName() {

	err := helper.InsertUser(&suite.mongoDB, suite.userForTest)
	assert.NoError(suite.T(), err)

	dbGateway, err := db.NewNoSQLDB(suite.mongoDB.URI(), suite.mongoDB.Name())
	assert.NoError(suite.T(), err)

	usr, err := dbGateway.FindUserByName(context.Background(), suite.userForTest.Name)

	suite.Nil(err)
	assertation.AssertThatUserExistsInDB(suite.T(), usr.Id.Hex(), &suite.mongoDB)
	assertation.AssertThatUserEqualWithouId(suite.T(), suite.userForTest, usr)
}

func (suite *DBSuite) TestFindUserByNameWithoutNoneExistentName() {

	err := helper.InsertUser(&suite.mongoDB, suite.userForTest)
	assert.NoError(suite.T(), err)

	dbGateway, err := db.NewNoSQLDB(suite.mongoDB.URI(), suite.mongoDB.Name())
	assert.NoError(suite.T(), err)

	usr, err := dbGateway.FindUserByName(context.Background(), "nonexistent name")

	suite.Error(err)
	suite.Empty(usr)
}

func (suite *DBSuite) TestDeleteUser() {

	err := helper.InsertUser(&suite.mongoDB, suite.userForTest)
	assert.NoError(suite.T(), err)

	dbGateway, err := db.NewNoSQLDB(suite.mongoDB.URI(), suite.mongoDB.Name())
	assert.NoError(suite.T(), err)

	err = dbGateway.DeleteUser(context.Background(), suite.userForTest.Name)

	suite.Nil(err)
	assertation.AssertThatUserDontExistsInDB(suite.T(), suite.userForTest.Name, &suite.mongoDB)
}

func (suite *DBSuite) TestDeleteDontExistentUser() {

	dbGateway, err := db.NewNoSQLDB(suite.mongoDB.URI(), suite.mongoDB.Name())
	assert.NoError(suite.T(), err)

	err = dbGateway.DeleteUser(context.Background(), suite.userForTest.Name)

	suite.Error(err)
}
