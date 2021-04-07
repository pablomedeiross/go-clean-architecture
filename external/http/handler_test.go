package http_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"strings"
	"testing"
	adapter "user-api/adapter/http"
	"user-api/external/http"
	"user-api/external/http/test"

	"github.com/stretchr/testify/assert"
)

const id_test = "Id"
const post = "POST"
const user_path = "/users"
const valid_create_user_request = "{" +
	"\"name\":\"exa\"," +
	"\"email\":\"meila\"," +
	"\"age\":12" +
	"}"

func TestCreateUser(t *testing.T) {

	controller := test.NewUserController(
		func(ctx context.Context, user adapter.User) (string, error) {
			return id_test, nil
		},
	)

	handler, _ := http.NewHandler(&controller)
	router := http.CreateEngineWithRoutes(&handler)

	responseRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(
		post,
		user_path,
		strings.NewReader(valid_create_user_request),
	)

	router.ServeHTTP(responseRecorder, req)
	assert.Equal(t, 201, responseRecorder.Code)
	assert.Equal(t, "http://example.com"+user_path+"/"+id_test, responseRecorder.Header().Get("Location"))
}

func TestCreateUserWithReturnBadRequestError(t *testing.T) {

	controller := test.NewUserController(nil)
	handler, _ := http.NewHandler(&controller)
	router := http.CreateEngineWithRoutes(&handler)

	responseRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(
		post,
		user_path,
		strings.NewReader(strings.Replace(valid_create_user_request, "}", "/", 2)),
	)

	errorExpected := http.Error{
		RequestPath:  user_path,
		RequestParms: "",
		RequestBody:  "{\"name\":\"\",\"email\":\"\",\"age\":0}",
		ErrorMsg:     "invalid character '/' after object key:value pair",
	}

	errorBytes, _ := json.Marshal(errorExpected)

	router.ServeHTTP(responseRecorder, req)
	assert.Equal(t, 400, responseRecorder.Code)
	assert.Equal(t, string(errorBytes), responseRecorder.Body.String())
}

func TestCreateUserWithReturnInternalError(t *testing.T) {

	controller := test.NewUserController(
		func(ctx context.Context, user adapter.User) (string, error) {
			return "", errors.New("Error")
		},
	)

	handler, _ := http.NewHandler(&controller)
	router := http.CreateEngineWithRoutes(&handler)

	responseRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(
		post,
		user_path,
		strings.NewReader(valid_create_user_request),
	)

	errorExpected := http.Error{
		RequestPath:  user_path,
		RequestParms: "",
		RequestBody:  valid_create_user_request,
		ErrorMsg:     "Error",
	}

	errorBytes, _ := json.Marshal(errorExpected)

	router.ServeHTTP(responseRecorder, req)
	assert.Equal(t, 500, responseRecorder.Code)
	assert.Equal(t, string(errorBytes), responseRecorder.Body.String())
}
