package handlers

import "github.com/valyala/fasthttp"

func StatusHandler(ctx *fasthttp.RequestCtx) {
	// notice that we may access MyHandler properties here - see h.foobar.
	RespondWithJSON(ctx, 200, "OK")
}
