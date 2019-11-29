package handlers

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type ErrorType string

const (
	MethodNotAllowed string = "method_not_allowed"
)

type errorHelper struct {
	Error       string `json:"error"`
	Description string `json:"description"`
}

type dataHelper struct {
	Data interface{} `json:"data"`
}

// RespondWithError - Returns an error response in JSON
func RespondWithError(ctx *fasthttp.RequestCtx, code int, errorType string, message string) {
	RespondWithJSON(ctx, code, &errorHelper{Error: errorType, Description: message})
}

// RespondWithJSON - Returns a JSON response
func RespondWithJSON(ctx *fasthttp.RequestCtx, code int, payload interface{}) {
	response, _ := json.Marshal(&dataHelper{Data: payload})

	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.SetStatusCode(code)
	ctx.Write(response)
}
