package utils

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(ctx *fasthttp.RequestCtx, data map[string]interface{}) {
	ctx.SetContentType("application/json")
	json.NewEncoder(ctx).Encode(data)
}
