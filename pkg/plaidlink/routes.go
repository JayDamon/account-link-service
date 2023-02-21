package plaidlink

import (
	"github.com/go-chi/chi/v5"
)

func AddRoutes(mux *chi.Mux, controller *Handler) {

	mux.Post("/v1/link/private-access-token", controller.CreatePrivateAccessToken)
	mux.Post("/v1/item/public-token", controller.CreateLinkToken)
}
