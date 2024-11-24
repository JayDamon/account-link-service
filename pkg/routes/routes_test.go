package routes

import (
	"context"
	"fmt"
	"github.com/factotum/moneymaker/account-link-service/pkg/config"
	"github.com/factotum/moneymaker/account-link-service/pkg/plaidlink"
	"github.com/go-chi/chi/v5"
	"github.com/jaydamon/moneymakerrabbit"
	"github.com/plaid/plaid-go/plaid"
	"github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestCreateRoutes_RoutesExist(t *testing.T) {

	os.Setenv("PLAID_CLIENT_ID", "test")
	os.Setenv("PLAID_SECRET", "test")
	os.Setenv("CLIENT_NAME", "test")
	os.Setenv("CLIENT_SECRET", "test")
	os.Setenv("REALM", "test")

	cnf := config.GetConfig()
	rabbitConnector := &TestConnector{}
	apiService := &TestApiService{}
	testHandler := plaidlink.NewHandler(cnf, apiService, rabbitConnector)

	routes := CreateRoutes(testHandler, cnf.KeyCloakConfig, false)
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

type TestConnector struct {
	body                 []interface{}
	headers              []map[string]interface{}
	contentType          []string
	queue                []string
	exchange             []string
	receiveMessagesCount int
	sendMessageCount     int
	failTimeCalled       int
}

func (conn *TestConnector) ReceiveMessages(queueName string, handler moneymakerrabbit.MessageHandlerFunc) {
	conn.receiveMessagesCount++
}

func (conn *TestConnector) SendMessage(body interface{}, headers map[string]interface{}, contentType string, queue string, exchange string) error {
	conn.body = append(conn.body, body)
	conn.headers = append(conn.headers, headers)
	conn.contentType = append(conn.contentType, contentType)
	conn.queue = append(conn.queue, queue)
	conn.exchange = append(conn.exchange, exchange)
	conn.sendMessageCount++

	if conn.sendMessageCount == conn.failTimeCalled {
		return fmt.Errorf("failing for %o test", conn.sendMessageCount)
	}

	return nil
}

func (conn *TestConnector) Close() {}

func (conn *TestConnector) DeclareExchange(exchangeName string) {}

func (conn *TestConnector) DeclareQueue(queueName string) *amqp091.Queue {
	return nil
}

func (conn *TestConnector) ReceiveMessagesFromExchange(exchangeName string, consumingQueueName string, handler moneymakerrabbit.MessageHandlerFunc) {
}

type TestApiService struct {
}

func (api *TestApiService) ItemPublicTokenExchange(ctx context.Context, tokenExchangeRequest *plaid.ItemPublicTokenExchangeRequest) (plaid.ItemPublicTokenExchangeResponse, *http.Response, error) {
	response := plaid.ItemPublicTokenExchangeResponse{}
	return response, nil, nil
}

func (api *TestApiService) RequestLinkToken(ctx context.Context, linkTokenRequest *plaid.LinkTokenCreateRequest) (plaid.LinkTokenCreateResponse, *http.Response, error) {
	response := plaid.LinkTokenCreateResponse{}
	return response, nil, nil
}
