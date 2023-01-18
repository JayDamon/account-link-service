package routes

import (
	"net/http"

	"github.com/factotum/moneymaker/plaid-integration/pkg/config"
	"github.com/factotum/moneymaker/plaid-integration/pkg/plaid"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jaydamon/moneymakergocloak"
)

func CreateRoutes(context *config.Context) http.Handler {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	keyCloakMiddleware := moneymakergocloak.NewKeyCloakMiddleWare(context.Config.KeyCloakConfig)
	router.Use(keyCloakMiddleware.VerifyToken)
	router.Use(middleware.Heartbeat("/ping"))

	plaid.AddRoutes(router, context)

	return router
}
