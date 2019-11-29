package handlers

import (
	"github.com/google/wire"
	"github.com/valyala/fasthttp"
)

var RouterSet = wire.NewSet(NewRouter)

type AppRequestHandler func(ctx *fasthttp.RequestCtx)

func statusHandler(ctx *fasthttp.RequestCtx) {
	// notice that we may access MyHandler properties here - see h.foobar.
	RespondWithJSON(ctx, 200, "OK")
}

func NewRouter() fasthttp.RequestHandler {
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		method := string(ctx.Method())
		switch string(ctx.Path()) {
		case "/status":
			if method == "GET" {
				statusHandler(ctx)
			} else {
				RespondWithError(ctx, 405, MethodNotAllowed, "")
			}
		case "/log":
			if method == "POST" {

			} else {
				RespondWithError(ctx, 405, MethodNotAllowed, "")
			}
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
	}

	return requestHandler
}
