package config

import "github.com/jaydamon/moneymakergocloak"

type Config struct {
	HostPort       string
	KeyCloakConfig *moneymakergocloak.KeyCloakConfig
}

func GetConfig() *Config {
	return &Config{
		HostPort: "8090",
		// HostPort: "3000",
		KeyCloakConfig: moneymakergocloak.NewKeyCloak(
			"http://localhost:8081/auth",
			"plaid-integration-service",
			"wQeV8pZwtBf9dIdKTGrqceyM3eeleokY",
			"moneymaker",
		),
	}
}
