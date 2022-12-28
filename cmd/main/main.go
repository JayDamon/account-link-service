package main

import (
	"log"

	"github.com/factotum/moneymaker/plaid-integration/pkg/app"
	"github.com/factotum/moneymaker/plaid-integration/pkg/config"
)

func main() {

	config := config.GetConfig()
	app := &app.App{}

	log.Print("Initializing app\n")

	app.Initialize(config)

	log.Printf("Starting service on port %s\n", config.HostPort)

	app.Run()
}
