package router

import (
	"chofu-line-bot/app/controllers"

	"github.com/gorilla/mux"
)

type AppRouter interface {
	SetAppRouting(router *mux.Router)
}

type appRouter struct {
	apc controllers.AppController
}

func NewAppRouter(apc controllers.AppController) AppRouter {
	return &appRouter{apc}
}

func (apr *appRouter) SetAppRouting(router *mux.Router) {
	router.HandleFunc("/", apr.apc.RootPage).Methods("GET")
}
