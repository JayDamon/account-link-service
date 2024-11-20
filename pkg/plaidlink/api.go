package plaidlink

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/factotum/moneymaker/account-link-service/pkg/config"
	"github.com/factotum/moneymaker/account-link-service/pkg/models"
	"github.com/factotum/moneymaker/account-link-service/pkg/users"
	tools "github.com/jaydamon/http-toolbox"
	"github.com/jaydamon/moneymakergocloak"
	"github.com/jaydamon/moneymakerrabbit"
	"github.com/plaid/plaid-go/plaid"
	"log"
	"net/http"
	"strings"
)

type Handler struct {
	config   *config.Config
	plaidApi ApiService
	rabbit   moneymakerrabbit.Connector
}

func NewHandler(config *config.Config, plaidApiService ApiService, rabbitConnection moneymakerrabbit.Connector) *Handler {
	return &Handler{
		config:   config,
		plaidApi: plaidApiService,
		rabbit:   rabbitConnection,
	}
}

func (handler *Handler) CreatePrivateAccessToken(w http.ResponseWriter, r *http.Request) {

	var publicToken models.PublicToken

	log.Print("Received request ", &r.Body)

	err := json.NewDecoder(r.Body).Decode(&publicToken)
	if err != nil {
		fmt.Println("Error decoding input body", err)
		tools.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ctx := context.Background()

	request := plaid.NewItemPublicTokenExchangeRequest(publicToken.PublicToken)
	exchangePublicTokenResp, _, err := handler.plaidApi.ItemPublicTokenExchange(ctx, request)
	if err != nil {
		fmt.Println("Error exchanging public token for private token", err)
		tools.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	accessToken := exchangePublicTokenResp.GetAccessToken()
	itemId := exchangePublicTokenResp.GetItemId()
	userId := moneymakergocloak.ExtractUserIdFromTokenFromRequest(w, r, handler.config.KeyCloakConfig)

	isNew := true
	pt := &models.PrivateToken{
		UserID:       &userId,
		PrivateToken: &accessToken,
		ItemId:       &itemId,
		IsNew:        &isNew,
	}

	bearerToken, err := moneymakergocloak.GetAuthorizationHeaderFromRequest(r)
	if err != nil {
		tools.RespondError(w, http.StatusUnauthorized, err.Error())
	}

	err = users.CreateAccountToken(handler.config, pt, bearerToken)
	if err != nil {
		fmt.Println("Error calling user service to create account token", err)
		tools.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	headers := make(map[string]interface{})
	headers["Authorization"] = bearerToken

	log.Printf("sending message to account_refresh queue %d", pt)
	err = handler.rabbit.SendMessage(pt, headers, "application/json", "account_refresh", "")
	if err != nil {
		fmt.Println("Error sending message", err)
		tools.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	tools.RespondNoBody(w, http.StatusCreated)
}

func (handler *Handler) CreateLinkToken(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	plaidConfig := handler.config.Plaid

	countryCodes := convertCountryCodes(strings.Split(plaidConfig.CountryCodes, ","))
	redirectURI := plaidConfig.RedirectUrl

	userId := moneymakergocloak.ExtractUserIdFromTokenFromRequest(w, r, handler.config.KeyCloakConfig)

	user := plaid.LinkTokenCreateRequestUser{
		ClientUserId: userId,
	}

	request := plaid.NewLinkTokenCreateRequest(
		"account-link-service", // Change to get this dynamically
		"en",
		countryCodes,
		user,
	)

	products := convertProducts(strings.Split(plaidConfig.Products, ","))
	request.SetProducts(products)

	if redirectURI != "" {
		request.SetRedirectUri(redirectURI)
	}

	linkTokenCreateResp, _, err := handler.plaidApi.RequestLinkToken(ctx, request)
	if err != nil {
		log.Print("Error retrieving LinkToken ", err)
		tools.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	linkToken := linkTokenCreateResp.GetLinkToken()
	tools.Respond(w, http.StatusOK, models.LinkToken{LinkToken: &linkToken})
}

func convertCountryCodes(countryCodeStrs []string) []plaid.CountryCode {
	var countryCodes []plaid.CountryCode

	for _, countryCodeStr := range countryCodeStrs {
		countryCodes = append(countryCodes, plaid.CountryCode(countryCodeStr))
	}

	return countryCodes
}

func convertProducts(productStrs []string) []plaid.Products {
	var products []plaid.Products

	for _, productStr := range productStrs {
		products = append(products, plaid.Products(productStr))
	}

	return products
}
