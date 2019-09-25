package core

import (
	"log"
	"net/http"
	"todo/core/driver/mysql"
	"todo/core/router"
	"todo/routes"

	"github.com/gorilla/mux"
)

type App struct{}

func (a *App) Run(addr string) {
	// Initialize DB Connection
	db := mysql.Mysql{}
	db.Initialize()

	// Core Router Instance
	rt := router.Router{}
	rt.Router = mux.NewRouter()

	// Get All Routes
	rts := routes.Routes{}
	rts.Router = &rt
	rts.InitializeRoutes()

	// Run HTTP Server
	log.Fatal(http.ListenAndServe(addr, rt.Router))
}
