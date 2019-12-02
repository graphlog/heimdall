package handlers

import (
	"github.com/google/wire"
	"github.com/graphlog/heimdall/pkg/services"
	"github.com/valyala/fasthttp"
	"net/http"
)

var RouterSet = wire.NewSet(NewRouter)

type AppRequestHandler func(ctx *fasthttp.RequestCtx)

func NewRouter(as *services.ApplicationService, ms *services.MessageService) fasthttp.RequestHandler {
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		method := string(ctx.Method())
		switch string(ctx.Path()) {
		case "/status":
			if method == "GET" {
				StatusHandler(ctx)
			} else {
				RespondWithError(ctx, 405, MethodNotAllowed, "")
			}
		case "/log":
			if method == "POST" {
				isValid := ValidateAPIKey(ctx, as)

				if !isValid {
					RespondWithError(ctx, http.StatusUnauthorized, UnAuthorized, "unauthorized request")
					return
				}

				SendLogToPlatform(ctx, ms)
			} else {
				RespondWithError(ctx, http.StatusMethodNotAllowed, MethodNotAllowed, "")
			}
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
	}

	return requestHandler
}
