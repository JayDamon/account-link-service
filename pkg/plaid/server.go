package plaid

import (
	"github.com/jaydamon/moneymakergocloak"
	"github.com/plaid/plaid-go/plaid"
)

var environments = map[string]plaid.Environment{
	"sandbox":     plaid.Sandbox,
	"development": plaid.Development,
	"production":  plaid.Production,
}

type PlaidConfig struct {
	Client       *plaid.APIClient
	Config       *plaid.Configuration
	Products     string
	CountryCodes string
	RedirectUrl  string
	Auth         *moneymakergocloak.KeyCloakConfig
}

func NewPlaidConfig(
	plaidClientId string,
	plaidSecret string,
	plaidEnv string,
	plaidProducts string,
	plaidCountryCodes string,
	plaidRedirectUri string,
	keyCloakConfig *moneymakergocloak.KeyCloakConfig) *PlaidConfig {

	config := plaid.NewConfiguration()
	config.AddDefaultHeader("PLAID-CLIENT-ID", plaidClientId)
	config.AddDefaultHeader("PLAID-SECRET", plaidSecret)
	config.UseEnvironment(environments[plaidEnv])
	// config.Debug = true

	client := plaid.NewAPIClient(config)

	return &PlaidConfig{
		Client:       client,
		Config:       config,
		Products:     plaidProducts,
		CountryCodes: plaidCountryCodes,
		RedirectUrl:  plaidRedirectUri,
		Auth:         keyCloakConfig,
	}
}
