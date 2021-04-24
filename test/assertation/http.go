package assertation

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"user-api/test/e2e/dto"

	"github.com/stretchr/testify/assert"
)

func AssertHttpPostWasRealized(t *testing.T, response http.Response, err error) {

	assert.NoError(t, err)
	assert.Equal(t, 201, response.StatusCode)
	assert.Contains(t, response.Header.Get(location), localhost_uri+user_path+"/")
}

func AssertHttpErrorEqual(t *testing.T, response http.Response, err error, expectedError dto.Error) {

	assert.NoError(t, err)
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	actualError := &dto.Error{}
	json.Unmarshal(bodyBytes, actualError)
	assert.Equal(t, expectedError, *actualError)
}
