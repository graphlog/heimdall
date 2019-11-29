package server

import (
	"github.com/google/wire"
	"github.com/valyala/fasthttp"
)

var ServerSet = wire.NewSet(CreateServer)

type AppServer struct {
	Port           string
	RequestHandler fasthttp.RequestHandler
}

func CreateServer(r fasthttp.RequestHandler) *AppServer {
	return &AppServer{
		Port:           ":9000",
		RequestHandler: r,
	}
}
