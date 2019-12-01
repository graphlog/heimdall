package handlers

import (
	"github.com/graphlog/heimdall/pkg/models"
	"github.com/graphlog/heimdall/pkg/services"
	"github.com/valyala/fasthttp"
)

func ValidateAPIKey(ctx *fasthttp.RequestCtx, as *services.ApplicationService) bool {
	key := string(ctx.Request.Header.Peek("GRAPHLOG_API_KEY"))
	secret := string(ctx.Request.Header.Peek("GRAPHLOG_API_SECRET"))

	if len(key) == 0 || len(secret) == 0 {
		return false
	}

	return as.Validate(&models.AuthorizationKeys{
		APIKey:    key,
		APISecret: secret,
	})
}
