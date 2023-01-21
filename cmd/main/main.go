package main

import (
	"log"

	"github.com/factotum/moneymaker/plaid-integration/pkg/app"
	"github.com/factotum/moneymaker/plaid-integration/pkg/config"
)

func main() {

	configuration := config.GetConfig()
	application := &app.App{}

	log.Print("Initializing application\n")

	application.Initialize(configuration)

	log.Printf("Starting service on port %s\n", configuration.HostPort)

	application.Run()
}
