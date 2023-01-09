package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/factotum/moneymaker/plaid-integration/pkg/config"
	"github.com/factotum/moneymaker/plaid-integration/pkg/routes"

	"github.com/go-chi/chi/v5"
)

type App struct {
	Router *chi.Mux
	Server *http.Server
}

func (a *App) Initialize(config *config.Config) {

	a.Server = &http.Server{
		Addr:    fmt.Sprintf(":%s", config.HostPort),
		Handler: routes.CreateRoutes(config),
	}
}

func (a *App) Run() {
	err := a.Server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
