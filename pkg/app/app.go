package app

import (
	"factotum/moneymaker/plaid-integration/pkg/config"
	"factotum/moneymaker/plaid-integration/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type App struct {
	Router *chi.Mux
	Server *http.Server
}

func (a *App) Initialize(config *config.Config) {

	a.Server = &http.Server{
		Addr:    fmt.Sprintf(":%s", config.HostPort),
		Handler: routes.CreateRoutes(a.Broker),
	}
}

func (a *App) Run() {
	err := a.Server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
