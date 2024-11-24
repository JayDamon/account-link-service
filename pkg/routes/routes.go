package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"net/http"

	"github.com/factotum/moneymaker/account-link-service/pkg/plaidlink"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jaydamon/moneymakergocloak"
)

func CreateRoutes(handler *plaidlink.Handler, keycloak *moneymakergocloak.Configuration, configureCors bool) http.Handler {
	router := chi.NewRouter()

	// TODO: remove cors config if behind api gateway. Setup in properties

	if configureCors {
		router.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
		}))
	}

	keyCloakMiddleware := moneymakergocloak.NewMiddleWare(keycloak)
	router.Use(keyCloakMiddleware.AuthorizeHttpRequest)
	router.Use(middleware.Heartbeat("/ping"))

	plaidlink.AddRoutes(router, handler)

	return router
}
