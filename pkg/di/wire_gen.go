// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/chocogem/slackbot-golang/pkg/api"
	"github.com/chocogem/slackbot-golang/pkg/api/handler"
	"github.com/chocogem/slackbot-golang/pkg/api/middleware"
	"github.com/chocogem/slackbot-golang/pkg/config"
	"github.com/chocogem/slackbot-golang/pkg/driver/client"
	"github.com/chocogem/slackbot-golang/pkg/driver/log"
	"github.com/chocogem/slackbot-golang/pkg/repository/adapter"
	adapter2 "github.com/chocogem/slackbot-golang/pkg/usecase/adapter"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*api.ServerHTTP, error) {
	apiClient := client.NewSlackApiClient()
	slackRepository, err := adapter.NewSlackRepository(cfg, apiClient)
	if err != nil {
		return nil, err
	}
	slackUseCase := adapter2.NewSlackUseCase(slackRepository)
	slackHandler := handler.NewSlackHandler(slackUseCase)
	handlers := &api.Handlers{
		SlackHandler: slackHandler,
	}
	logger := log.NewLogrusLogger(cfg)
	logrusImplement := log.NewImplementLogrusLogger(logger)
	errorHandler := middleware.NewErrorHandler(logrusImplement)
	middlewares := &api.Middlewares{
		ErrorHandler: errorHandler,
	}
	serverHTTP := api.NewServerHTTP(handlers, middlewares)
	return serverHTTP, nil
}