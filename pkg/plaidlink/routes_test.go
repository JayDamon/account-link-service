package plaidlink

import (
	"github.com/factotum/moneymaker/account-link-service/pkg/config"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestAddRoutes(t *testing.T) {

	configuration := &config.Config{}
	testHandler := NewHandler(configuration)

	router := chi.NewRouter()

	AddRoutes(router, testHandler)

	assert.NotNil(t, router)

	routeExists(t, router, "/v1/link/private-access-token")
	routeExists(t, router, "/v1/item/public-token")
}

func routeExists(t *testing.T, routes chi.Router, routeToValidate string) {
	found := false

	_ = chi.Walk(routes, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if route == routeToValidate {
			found = true
		}
		return nil
	})
	assert.True(t, found, "route not found %s", routeToValidate)
}
