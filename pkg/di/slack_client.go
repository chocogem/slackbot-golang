package di

import (
	"github.com/google/wire"
	"github.com/chocogem/slackbot-golang/pkg/driver/client"
)

var SlackClientSet = wire.NewSet(
	client.NewSlackApiClient,
)
