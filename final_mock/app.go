package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialise() {
	a.Router = mux.NewRouter()

	a.Router.HandleFunc("/blueprints", GetBlueprints).Methods("GET")
	a.Router.HandleFunc("/blueprint", AddBlueprint).Methods("POST")
	a.Router.HandleFunc("/blueprints/{id}", GetBlueprint).Methods("GET")

}

func (a *App) Run(addr string) {
	http.ListenAndServe(":9090", a.Router)
}
