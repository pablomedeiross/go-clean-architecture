package http

import (
	"github.com/gin-gonic/gin"
)

const base_path = "/users"

func CreateEngineWithRoutes(httpHandler *Handler) *gin.Engine {

	engine := gin.New()
	engine.Use(gin.Logger())
	engine.POST(base_path, httpHandler.Post)
	engine.DELETE(base_path+"/:id", httpHandler.Delete)
	return engine
}
