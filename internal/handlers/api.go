package handlers

import (
	"fmt"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/amit-chourey/go_temp/inetrnal/middleware"
)

func Handler(r *chi.Mux){
	r.use(chimiddle.StrpSlashes)
	r.Route("/account", func(router chi.Router){
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})
}

