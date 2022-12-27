package main

import (
	"factotum/moneymaker/plaid-integration/pkg/app"
	"factotum/moneymaker/plaid-integration/pkg/config"
	"log"
)

func main() {

	config := config.GetConfig()
	app := &app.App{}

	log.Printf("Starting service on port %s\n", config.HostPort)

	app.Initialize(config)

	app.Run()
}
