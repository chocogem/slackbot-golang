package api

import (
	"fmt"
	"github.com/chocogem/slackbot-golang/pkg/api/handler"
	"github.com/chocogem/slackbot-golang/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

type Handlers struct {
	SlackHandler *handler.SlackHandler
}
type Middlewares struct {
	ErrorHandler   *middleware.ErrorHandler
}

func NewServerHTTP(handlers *Handlers, middlewares *Middlewares) *ServerHTTP {
	engine := gin.New()
	engine.Use(middlewares.ErrorHandler.Handler())

	api := engine.Group("/slack")
	api.POST("/hello", handlers.SlackHandler.SendHello)
	api.POST("/sendCustomMessage", handlers.SlackHandler.SendCustomMessage)
	return &ServerHTTP{engine: engine}
}
func (sh *ServerHTTP) Start(appPort string) {
	sh.engine.Run(fmt.Sprintf(":%s", appPort))
}
