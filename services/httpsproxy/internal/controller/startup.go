package controller

import (
	"github.com/julienschmidt/httprouter"
	"github.com/kuritka/gsvc/services/httpsproxy/internal/proxy"
)

var router *httprouter.Router

func init() {
	router = httprouter.New()
}

func Startup(defaultHost string) {
	px := proxy.NewHttpsProxy(defaultHost)

	router.GET("/", px.Handle)

	go px.Listen()
}

func Router() *httprouter.Router {
	return router
}
