package di

import (
	repoAdapter "github.com/chocogem/slackbot-golang/pkg/repository/adapter"
	usecaseAdapter "github.com/chocogem/slackbot-golang/pkg/usecase/adapter"
	"github.com/chocogem/slackbot-golang/pkg/api/handler"
	"github.com/google/wire"
)

var SlackSet = wire.NewSet(
	repoAdapter.NewSlackRepository,
	usecaseAdapter.NewSlackUseCase,
	handler.NewSlackHandler,
)
