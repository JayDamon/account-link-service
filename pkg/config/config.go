package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jaydamon/moneymakergocloak"
	"github.com/joho/godotenv"
)

type Config struct {
	HostPort       string
	UserServiceUrl string
	KeyCloakConfig *moneymakergocloak.KeyCloakConfig
	Plaid          *PlaidConfig
}

func GetConfig() *Config {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error when loading environment variables from .env file %w", err)
	}

	hostPort := getOrDefault("HOST_PORT", "3000")

	userServiceUrl := getOrDefault("USER_SERVICE_URL", "http://localhost:8091")

	keycloakIssuerUri := getOrDefault("ISSUER_URI", "http://keycloak:8081/auth")
	keycloakClientName := getOrDefault("CLIENT_NAME", "account-link-service")
	keycloakClientSecret := getOrDefault("CLIENT_SECRET", "wQeV8pZwtBf9dIdKTGrqceyM3eeleokY")
	keycloakRealm := getOrDefault("REALM", "moneymaker")
	keyCloakDebugActive := getOrDefaultBool("DEBUG_ACTIVE", false)

	keyCloakConfig := moneymakergocloak.NewKeyCloak(
		keycloakIssuerUri,
		keycloakClientName,
		keycloakClientSecret,
		keycloakRealm,
		keyCloakDebugActive,
	)

	return &Config{
		HostPort:       hostPort,
		UserServiceUrl: userServiceUrl,
		KeyCloakConfig: keyCloakConfig,
		Plaid:          getPlaidConfig(keyCloakConfig),
	}
}

func getPlaidConfig(auth *moneymakergocloak.KeyCloakConfig) *PlaidConfig {

	PLAID_CLIENT_ID := getOrExit("PLAID_CLIENT_ID")
	PLAID_SECRET := getOrExit("PLAID_SECRET")
	PLAID_ENV := getOrDefault("PLAID_ENV", "sandbox")
	PLAID_PRODUCTS := getOrDefault("PLAID_PRODUCTS", "transactions")
	PLAID_COUNTRY_CODES := getOrDefault("PLAID_COUNTRY_CODES", "US")
	PLAID_REDIRECT_URI := getOrDefault("PLAID_REDIRECT_URI", "")

	return newPlaidConfig(
		PLAID_CLIENT_ID,
		PLAID_SECRET,
		PLAID_ENV,
		PLAID_PRODUCTS,
		PLAID_COUNTRY_CODES,
		PLAID_REDIRECT_URI,
		auth,
	)
}

func getOrDefault(envVar string, defaultVal string) string {
	val := os.Getenv(envVar)
	if val == "" {
		return defaultVal
	}
	return val
}

func getOrDefaultBool(envVar string, defaultVal bool) bool {
	val := os.Getenv(envVar)
	var returnVal = defaultVal
	if val == "true" {
		returnVal = true
	} else if val == "false" {
		returnVal = false
	}

	return returnVal
}

func getOrExit(envVar string) string {
	val := os.Getenv(envVar)
	if val == "" {
		log.Fatal(fmt.Printf("%s is not set. Make sure to fill out the .env file", envVar))
	}
	return val
}
