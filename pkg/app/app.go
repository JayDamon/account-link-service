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
	Router  *chi.Mux
	Server  *http.Server
	Context *config.Context
}

func (a *App) Initialize(configuration *config.Config) {

	a.Context = &config.Context{
		Config: configuration,
	}

	a.Server = &http.Server{
		Addr:    fmt.Sprintf(":%s", configuration.HostPort),
		Handler: routes.CreateRoutes(a.Context),
	}
}

func (a *App) Run() {
	err := a.Server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
