package plaid

import (
	"github.com/factotum/moneymaker/plaid-integration/pkg/config"
	"github.com/go-chi/chi/v5"
)

func AddRoutes(mux *chi.Mux, context *config.Context) {

	plaidContext := &Context{
		context: context,
	}

	mux.Post("/v1/link/private-access-token", plaidContext.CreatePrivateAccessToken)
	mux.Post("/v1/item/public-token", plaidContext.CreateLinkToken)
}
