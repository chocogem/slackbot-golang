//go:build wireinject
// +build wireinject

package di

import (
	"github.com/chocogem/slackbot-golang/pkg/api"
	"github.com/chocogem/slackbot-golang/pkg/config"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*api.ServerHTTP, error) {
	wire.Build(SlackSet,HTTPSet,LogSet,SlackClientSet)
	return &api.ServerHTTP{}, nil
}
