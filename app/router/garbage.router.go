package router

import (
	"chofu-line-bot/app/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

type GarbageRouter interface {
	SetGarbageRouting(router *mux.Router)
}

type garbageRouter struct {
	gc controllers.GarbageDayController
}

func NewGarbageRouter(gc controllers.GarbageDayController) GarbageRouter {
	return &garbageRouter{gc}
}

func (gr *garbageRouter) SetGarbageRouting(router *mux.Router) {
	router.Handle("/callback", http.HandlerFunc(gr.gc.Callback)).Methods("POST")
}
