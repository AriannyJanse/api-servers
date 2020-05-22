package server

import (
	"api-server/controllers"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

type api struct {
	router *fasthttprouter.Router
}

type Server interface {
	Router() *fasthttprouter.Router
}

func (a *api) Router() *fasthttprouter.Router {
	return a.router
}

func New() Server {
	a := &api{}

	r := fasthttprouter.New()
	r.GET("/servers", controllers.GetServers)
	r.GET("/servers/:domain", controllers.GetServerByDomain)

	a.router = r
	return a
}

var (
	corsAllowHeaders     = "authorization"
	corsAllowMethods     = "HEAD,GET"
	corsAllowOrigin      = "*"
	corsAllowCredentials = "true"
)

func CORS(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {

		ctx.Response.Header.Set("Access-Control-Allow-Credentials", corsAllowCredentials)
		ctx.Response.Header.Set("Access-Control-Allow-Headers", corsAllowHeaders)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", corsAllowMethods)
		ctx.Response.Header.Set("Access-Control-Allow-Origin", corsAllowOrigin)

		next(ctx)
	}
}
