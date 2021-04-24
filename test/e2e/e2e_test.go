package e2e_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"user-api/db/inmemory"
	"user-api/external/configuration"
	"user-api/test/assertation"
	"user-api/test/e2e/dto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	id_test       = "Id"
	post          = "POST"
	user_path     = "/users"
	content_type  = "application/json"
	localhost_uri = "http://localhost:8080"
)

type E2ESuite struct {
	suite.Suite
	mongoDB     inmemory.InMemoryMongoDB
	application configuration.AppStarter
}

func (suite *E2ESuite) SetupSuite() {

	suite.mongoDB = *inmemory.NewInMemoryMongoDB()
	suite.mongoDB.Start()

	starter, _ := configuration.NewAppStarter("local")
	suite.application = *starter
	go suite.application.Start()
}

func (suite *E2ESuite) TearDownSuite() {
	suite.mongoDB.Stop()
}

func TestE2ESuite(t *testing.T) {
	suite.Run(t, new(E2ESuite))
}

func (suite *E2ESuite) TestCreateUser() {

	newUser := dto.RequestNewUser{Name: "name1", Email: "email@gmail.com", Age: 12}
	jsonRequest, _ := json.Marshal(newUser)
	response, err := sendPostToCreateUser(jsonRequest)
	assertation.AssertThatUserWasCreated(suite.T(), *response, err)
}

func (suite *E2ESuite) TestCreateUserThatAlreadyExists() {

	newUser := dto.RequestNewUser{Name: "name2", Email: "email@uol.com", Age: 12}
	jsonRequest, _ := json.Marshal(newUser)

	expectedError := dto.Error{
		RequestPath:  user_path,
		RequestParms: "",
		RequestBody:  string(jsonRequest),
		ErrorMsg: "Couldn't create new user, usecase returning error: " +
			"The user already existis, user's name: " +
			newUser.Name,
	}

	response, err := sendPostToCreateUser(jsonRequest)
	assertation.AssertThatUserWasCreated(suite.T(), *response, err)

	response, err = sendPostToCreateUser(jsonRequest)
	assertation.AssertHttpErrorEqual(suite.T(), *response, err, expectedError)
}

func (suite *E2ESuite) TestDeleteUser() {

	newUser := dto.RequestNewUser{Name: "name4", Email: "email@hotmail.com", Age: 12}
	jsonRequest, _ := json.Marshal(newUser)
	response, err := sendPostToCreateUser(jsonRequest)

	assertation.AssertThatUserWasCreated(suite.T(), *response, err)

	request, _ := http.NewRequest(http.MethodDelete, localhost_uri+user_path+"/"+newUser.Name, nil)
	client := &http.Client{}
	response, err = client.Do(request)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 204, response.StatusCode)
	assertation.AssertThatUserDontExistsInDB(suite.T(), newUser.Name, &suite.mongoDB)
}

func (suite *E2ESuite) TestDeleteNoExistentUser() {

	expectedError := dto.Error{
		RequestPath:  user_path + "/:id",
		RequestParms: "",
		RequestBody:  "",
		ErrorMsg: "Couldn't remove user, usecase returning error: " +
			"Error to try remove a user in RemoveUserInteractor: " +
			"Error to delete user in repository: " +
			"no exists a user with this name in database to delection",
	}

	request, _ := http.NewRequest(http.MethodDelete, localhost_uri+user_path+"/"+"noexistentuser", nil)
	client := &http.Client{}
	response, err := client.Do(request)

	assertation.AssertHttpErrorEqual(suite.T(), *response, err, expectedError)
}

func sendPostToCreateUser(jsonRequest []byte) (*http.Response, error) {
	return http.Post(localhost_uri+user_path, content_type, bytes.NewReader(jsonRequest))
}
