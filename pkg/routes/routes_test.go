package routes

import (
	"github.com/factotum/moneymaker/account-link-service/pkg/config"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestCreateRoutes_RoutesExist(t *testing.T) {

	os.Setenv("PLAID_CLIENT_ID", "test")
	os.Setenv("PLAID_SECRET", "test")

	keyCloakConfig := config.GetConfig()

	context := config.Context{
		Config: keyCloakConfig,
	}

	routes := CreateRoutes(&context)
	chiRoutes := routes.(chi.Router)

	assert.NotNil(t, chiRoutes)

	routeExists(t, chiRoutes, "/v1/link/private-access-token")
	routeExists(t, chiRoutes, "/v1/item/public-token")
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
