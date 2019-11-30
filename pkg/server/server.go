package server

import (
	"fmt"
	"github.com/google/wire"
	"github.com/graphlog/heimdall/pkg/config"
	"github.com/valyala/fasthttp"
)

var ServerSet = wire.NewSet(config.NewConfig, CreateServer)

type AppServer struct {
	Port           string
	RequestHandler fasthttp.RequestHandler
}

func CreateServer(c *config.Config, r fasthttp.RequestHandler) *AppServer {
	return &AppServer{
		Port:           fmt.Sprintf(":%s", c.Server.Port),
		RequestHandler: r,
	}
}
