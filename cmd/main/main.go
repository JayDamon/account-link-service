package main

import (
	"log"

	"github.com/factotum/moneymaker/account-link-service/pkg/app"
)

func main() {

	application := app.NewApplication()

	log.Print("Initializing application\n")

	application.Initialize()

	defer application.RabbitConnection.Close()

	application.Run()
}
