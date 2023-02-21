package app

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"

	"github.com/factotum/moneymaker/account-link-service/pkg/config"
	"github.com/factotum/moneymaker/account-link-service/pkg/plaidlink"
	"github.com/factotum/moneymaker/account-link-service/pkg/routes"
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

	handler := plaidlink.NewHandler(configuration)

	a.Server = &http.Server{
		Addr:    fmt.Sprintf(":%s", configuration.HostPort),
		Handler: routes.CreateRoutes(handler, configuration.KeyCloakConfig),
	}
}

func (a *App) Run() {
	err := a.Server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
