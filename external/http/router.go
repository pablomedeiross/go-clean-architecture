package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(httpHandler *Handler) *gin.Engine {

	r := gin.Default()
	r.POST("/users", httpHandler.Post)
	return r

}
