package plaidlink

import (
	"fmt"
	"github.com/factotum/moneymaker/account-link-service/pkg/config"
	"github.com/go-chi/chi/v5"
	"github.com/jaydamon/moneymakerrabbit"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestAddRoutes(t *testing.T) {

	configuration := &config.Config{}
	rabbitConnector := &TestConnector{}
	testHandler := NewHandler(configuration, rabbitConnector)

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
