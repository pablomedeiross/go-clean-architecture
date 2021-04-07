package e2etest

import (
	"user-api/external/configuration"
	"user-api/external/db/memory"

	"github.com/stretchr/testify/suite"
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
	suite.application.Start()
}

func (suite *E2ESuite) TearDownSuite() {
	suite.mongoDB.Stop()
}

// func (suite *E2ESuite) TestCool() {
// 	assert.Equal(suite.T(), 5, 5)
// }

// func TestE2ESuite(t *testing.T) {
// 	suite.Run(t, new(E2ESuite))
// }
