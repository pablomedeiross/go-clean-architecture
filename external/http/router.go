package http

import (
	"github.com/gin-gonic/gin"
)

func CreateEngineWithRoutes(httpHandler *Handler) *gin.Engine {

	engine := gin.New()
	engine.Use(gin.Logger())
	engine.POST("/users", httpHandler.Post)
	return engine
}
