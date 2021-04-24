package http

import (
	"errors"
	adapter "user-api/adapter/http"

	"github.com/gin-gonic/gin"
)

const (
	controller_nil_create_http_handler = "Error to create HttpHanlder, controller received is nul"
)

type Handler struct {
	controller adapter.HttpController
}

func NewHandler(controller *adapter.HttpController) (Handler, error) {

	if controller == nil {
		return Handler{}, errors.New(controller_nil_create_http_handler)
	}

	return Handler{controller: *controller}, nil
}

func (handler *Handler) Post(context *gin.Context) {

	var usr *adapter.User = &adapter.User{}
	err := context.BindJSON(usr)

	if err != nil {
		createBadRequestResponse(err, usr, context)
		return
	}

	id, err :=
		handler.
			controller.
			CreateUser(context, *usr)

	if err != nil {
		createInternalServerErrorResponse(err, usr, context)
		return
	}

	addLocationHeader(id, context)
}

func (handler *Handler) Delete(ctx *gin.Context) {

	param := ctx.Param("id")

	err := handler.
		controller.
		RemoveUser(ctx, param)

	if err != nil {
		createInternalServerErrorResponse(err, nil, ctx)
	}

	ctx.Status(204)
}
