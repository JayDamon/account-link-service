package config

import (
	"fmt"
	"github.com/jaydamon/moneymakerplaid"
	"os"

	"github.com/jaydamon/moneymakergocloak"
	"github.com/joho/godotenv"
)

type Config struct {
	HostPort       string
	UserServiceUrl string
	KeyCloakConfig *moneymakergocloak.Configuration
	Plaid          *moneymakerplaid.Configuration
}

func GetConfig() *Config {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error when loading environment variables from .env file %w", err)
	}

	hostPort := getOrDefault("HOST_PORT", "3000")

	userServiceUrl := getOrDefault("USER_SERVICE_URL", "http://localhost:8091")

	keyCloakConfig := moneymakergocloak.NewConfiguration()

	return &Config{
		HostPort:       hostPort,
		UserServiceUrl: userServiceUrl,
		KeyCloakConfig: keyCloakConfig,
		Plaid:          moneymakerplaid.NewConfiguration(),
	}
}

func getOrDefault(envVar string, defaultVal string) string {
	val := os.Getenv(envVar)
	if val == "" {
		return defaultVal
	}
	return val
}
