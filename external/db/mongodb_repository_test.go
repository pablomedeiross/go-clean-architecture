package db_test

import (
	"context"
	"testing"
	adapter "user-api/adapter/db"
	"user-api/external/db"
	"user-api/external/db/test/assertation"
	"user-api/helper"

	test_helper "user-api/external/db/test/helper"

	"github.com/stretchr/testify/suite"
)

type DBSuite struct {
	suite.Suite
	mongoDB     helper.InMemoryMongoDB
	userForTest adapter.User
}

func (suite *DBSuite) SetupSuite() {

	suite.mongoDB = *helper.NewInMemoryMongoDB()
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
	id, err := dbGateway.SaveUser(context.Background(), suite.userForTest)

	assertation.AssertThatUserExistsInDB(suite.T(), id, &suite.mongoDB)
	suite.Nil(err)
}

func (suite *DBSuite) TestFindUserByName() {

	err := test_helper.InsertUser(&suite.mongoDB, suite.userForTest)

	dbGateway, err := db.NewNoSQLDB(suite.mongoDB.URI(), suite.mongoDB.Name())
	usr, err := dbGateway.FindUserByName(context.Background(), suite.userForTest.Name)

	suite.Nil(err)
	assertation.AssertThatUserExistsInDB(suite.T(), usr.Id, &suite.mongoDB)
	assertation.AssertThatUserEqualWithouId(suite.T(), suite.userForTest, usr)
}

func (suite *DBSuite) TestFindUserByNameWithoutNoneExistentName() {

	err := test_helper.InsertUser(&suite.mongoDB, suite.userForTest)

	dbGateway, err := db.NewNoSQLDB(suite.mongoDB.URI(), suite.mongoDB.Name())
	usr, err := dbGateway.FindUserByName(context.Background(), "nonexistent name")

	suite.Error(err)
	suite.Empty(usr)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(DBSuite))
}
