package app

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/jaydamon/moneymakerrabbit"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"net/http"

	"github.com/factotum/moneymaker/account-link-service/pkg/config"
	"github.com/factotum/moneymaker/account-link-service/pkg/plaidlink"
	"github.com/factotum/moneymaker/account-link-service/pkg/routes"
)

type App struct {
	Router           *chi.Mux
	Server           *http.Server
	RabbitConnection moneymakerrabbit.Connector
	RabbitChannel    *amqp091.Channel
	Config           *config.Config
}

func NewApplication() *App {
	return &App{
		Config: config.GetConfig(),
	}
}

func (a *App) Initialize() {

	a.RabbitConnection = a.Config.Rabbit.Connect()
	plaidApi := plaidlink.NewApiService(a.Config.Plaid.Config)

	handler := plaidlink.NewHandler(a.Config, plaidApi, a.RabbitConnection)
	a.RabbitConnection.DeclareQueue("account_refresh")

	a.Server = &http.Server{
		Addr:    fmt.Sprintf(":%s", a.Config.HostPort),
		Handler: routes.CreateRoutes(handler, a.Config.KeyCloakConfig),
	}

}

func (a *App) Run() {
	appName := a.Config.ApplicationName
	if appName == "" {
		appName = "unnamed service"
	}
	log.Printf("Starting \"%s\" service on port %s\n", appName, a.Config.HostPort)
	err := a.Server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
