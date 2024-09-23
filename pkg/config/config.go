package config

import (
	"fmt"
	"github.com/jaydamon/moneymakerplaid"
	"github.com/jaydamon/moneymakerrabbit"
	"os"

	"github.com/jaydamon/moneymakergocloak"
	"github.com/joho/godotenv"
)

type Config struct {
	HostPort        string
	ApplicationName string
	UserServiceUrl  string
	KeyCloakConfig  *moneymakergocloak.Configuration
	Plaid           *moneymakerplaid.Configuration
	Rabbit          *moneymakerrabbit.Configuration
}

func GetConfig() *Config {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error when loading environment variables from .env file %w", err)
	}

	hostPort := getOrDefault("HOST_PORT", "3000")
	applicationName := getOrDefault("APPLICATION_NAME", "")

	userServiceUrl := getOrDefault("USER_SERVICE_URL", "http://localhost:8091")

	return &Config{
		HostPort:        hostPort,
		ApplicationName: applicationName,
		UserServiceUrl:  userServiceUrl,
		KeyCloakConfig:  moneymakergocloak.NewConfiguration(),
		Plaid:           moneymakerplaid.NewConfiguration(),
		Rabbit:          moneymakerrabbit.NewConfiguration(),
	}
}

func getOrDefault(envVar string, defaultVal string) string {
	val := os.Getenv(envVar)
	if val == "" {
		return defaultVal
	}
	return val
}
