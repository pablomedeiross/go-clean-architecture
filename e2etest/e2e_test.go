package e2etest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"
	"user-api/e2etest/assertation"
	"user-api/e2etest/dto"
	"user-api/external/configuration"
	"user-api/external/db/memory"

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
	mongoDB     memory.InMemoryMongoDB
	application configuration.AppStarter
}

func (suite *E2ESuite) SetupSuite() {

	suite.mongoDB = *memory.NewInMemoryMongoDB()
	suite.mongoDB.Start()

	starter, _ := configuration.NewAppStarter("local")
	suite.application = *starter
	go suite.application.Start()
	time.Sleep(5000)
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

	newUser := dto.RequestNewUser{Name: "name2", Email: "email@gmail.com", Age: 12}
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
	assertation.AssertThatUserAlreadyExists(suite.T(), *response, err, expectedError)
}

func sendPostToCreateUser(jsonRequest []byte) (*http.Response, error) {
	return http.Post(localhost_uri+user_path, content_type, bytes.NewReader(jsonRequest))
}
