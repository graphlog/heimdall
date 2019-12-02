package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/graphlog/heimdall/pkg/models"
	"github.com/graphlog/heimdall/pkg/services"
	"github.com/valyala/fasthttp"
)

func SendLogToPlatform(ctx *fasthttp.RequestCtx, ms *services.MessageService) {
	var log models.Log
	reader := bytes.NewReader(ctx.PostBody())
	decode := json.NewDecoder(reader)

	if err := decode.Decode(&log); err != nil {
		RespondWithError(ctx, http.StatusBadRequest, InvalidRequest, "invalid message in log")
		return
	}

	err := ms.SendLogs(&log)

	if err != nil {
		RespondWithError(ctx, http.StatusInternalServerError, InvalidRequest, "failed to send message to platform")
		return
	}

	RespondWithJSON(ctx, http.StatusOK, "done")
	return
}
