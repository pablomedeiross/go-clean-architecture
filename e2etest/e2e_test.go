package e2etest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"user-api/e2etest/dto"
	"user-api/external/configuration"
	"user-api/external/db/memory"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	id_test       = "Id"
	post          = "POST"
	user_path     = "/users"
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
}

func (suite *E2ESuite) TearDownSuite() {
	suite.mongoDB.Stop()
}

func TestE2ESuite(t *testing.T) {
	suite.Run(t, new(E2ESuite))
}

func (suite *E2ESuite) TestCreateUser() {

	request := dto.RequestNewUser{Name: "name", Email: "email@gmail.com", Age: 12}
	jsonRequest, _ := json.Marshal(request)

	response, err := http.Post(localhost_uri+user_path, "application/json", bytes.NewReader(jsonRequest))

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 201, response.StatusCode)
	assert.Contains(suite.T(), response.Header.Get("Location"), localhost_uri+user_path+"/")
}
