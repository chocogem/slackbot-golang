package di

import (
	"github.com/chocogem/slackbot-golang/pkg/api"
	"github.com/google/wire"
	"github.com/chocogem/slackbot-golang/pkg/api/middleware"
)

var HTTPSet = wire.NewSet(
	api.NewServerHTTP,
	
	middleware.NewErrorHandler,
	wire.Struct(new(api.Middlewares), "*"),
	wire.Struct(new(api.Handlers), "*"),
)
