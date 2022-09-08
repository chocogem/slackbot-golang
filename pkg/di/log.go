package di

import (
	"github.com/google/wire"
	"github.com/chocogem/slackbot-golang/pkg/driver/log"
)

var LogSet = wire.NewSet(
	log.NewImplementLogrusLogger, 
	log.NewLogrusLogger,
	wire.Bind(new(log.Logger),new(*log.LogrusImplement)), 
	
)
