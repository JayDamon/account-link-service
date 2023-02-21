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
	Router *chi.Mux
	Server *http.Server
	Config *config.Config
}

func (a *App) Initialize() {

	handler := plaidlink.NewHandler(a.Config)

	a.Server = &http.Server{
		Addr:    fmt.Sprintf(":%s", a.Config.HostPort),
		Handler: routes.CreateRoutes(handler, a.Config.KeyCloakConfig),
	}
}

func (a *App) Run() {
	err := a.Server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
