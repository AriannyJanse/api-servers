package server

import (
	"api-server/controllers"

	"github.com/buaazp/fasthttprouter"
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
