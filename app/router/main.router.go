package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type MainRouter interface {
	setupRouting() *mux.Router
	StartWebServer() error
}

type mainRouter struct {
	appR     AppRouter
	garbageR GarbageRouter
}

func NewMainRouter(appR AppRouter, garbageR GarbageRouter) MainRouter {
	return &mainRouter{appR, garbageR}
}

func (mainRouter *mainRouter) setupRouting() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	mainRouter.appR.SetAppRouting(router)
	mainRouter.garbageR.SetGarbageRouting(router)

	return router
}

func (mainRouter *mainRouter) StartWebServer() error {

	return http.ListenAndServe(":8080", mainRouter.setupRouting())

}
