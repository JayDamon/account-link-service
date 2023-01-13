package routes

import (
	"net/http"

	"github.com/factotum/moneymaker/plaid-integration/pkg/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jaydamon/moneymakergocloak"
)

func CreateRoutes(config *config.Config) http.Handler {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	keyCloakMiddleware := moneymakergocloak.NewKeyCloakMiddleWare(config.KeyCloakConfig)
	router.Use(keyCloakMiddleware.VerifyToken)
	router.Use(middleware.Heartbeat("/ping"))

	addRoutes(router, config)

	return router
}

func addRoutes(mux *chi.Mux, config *config.Config) {
	mux.Post("/api/v1/link/private-access-token", config.Plaid.CreatePrivateAccessToken)
	mux.Post("/api/v1/item/public-token", config.Plaid.CreateLinkToken)
}
