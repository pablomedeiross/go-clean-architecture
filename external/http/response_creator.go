package http

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	http_prefix = "http://"
	bar         = "/"
	location    = "Location"
)

//addLocationHeader add location header in actual gin.Context
func addLocationHeader(idResource string, context *gin.Context) {

	urlString :=
		http_prefix +
			context.Request.Host +
			context.FullPath() +
			bar +
			idResource

	context.Header(location, urlString)
	context.JSON(http.StatusCreated, nil)
}

// createBadRequestResponse create a new error response with status 400 in actual gin.Context
func createBadRequestResponse(err error, requestBody interface{}, context *gin.Context) {
	createErrorResponse(err, http.StatusBadRequest, requestBody, context)
}

// createInternalServerErrorResponse create new error response with status 500 in actual gin.Context
func createInternalServerErrorResponse(err error, requestBody interface{}, context *gin.Context) {
	createErrorResponse(err, http.StatusInternalServerError, requestBody, context)
}

// createErrorResponse create a new error response in actual gin.Context
func createErrorResponse(err error, status int, requestBody interface{}, context *gin.Context) {

	json, marshalError := json.Marshal(requestBody)

	if marshalError != nil {
		err = marshalError
		json = []byte{}
	}

	context.JSON(
		status,
		Error{
			RequestPath:  context.FullPath(),
			RequestParms: context.Request.URL.RawQuery,
			RequestBody:  string(json),
			ErrorMsg:     err.Error(),
		},
	)
}
