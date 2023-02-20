package routes

import (
	"github.com/go-chi/chi/v5"
	"net/http"

	"github.com/factotum/moneymaker/account-link-service/pkg/config"
	"github.com/factotum/moneymaker/account-link-service/pkg/plaid"
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
