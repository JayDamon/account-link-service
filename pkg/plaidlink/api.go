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
	"github.com/plaid/plaid-go/plaid"
	"log"
	"net/http"
	"strings"
)

type Handler struct {
	config *config.Config
}

func NewHandler(appConfig *config.Config) *Handler {
	return &Handler{
		config: appConfig,
	}
}

func (handler *Handler) CreatePrivateAccessToken(w http.ResponseWriter, r *http.Request) {

	var publicToken models.PublicToken

	plaidConfig := handler.config.Plaid

	log.Print("Recieved request ", &r.Body)

	err := json.NewDecoder(r.Body).Decode(&publicToken)
	if err != nil {
		fmt.Println("Error decoding input body", err)
		tools.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ctx := context.Background()

	client := plaidConfig.Client
	exchangePublicTokenResp, _, err := client.PlaidApi.ItemPublicTokenExchange(ctx).ItemPublicTokenExchangeRequest(
		*plaid.NewItemPublicTokenExchangeRequest(publicToken.PublicToken),
	).Execute()
	if err != nil {
		fmt.Println("Error exchanging public token for private token", err)
		tools.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	accessToken := exchangePublicTokenResp.GetAccessToken()
	itemId := exchangePublicTokenResp.GetItemId()
	userId := moneymakergocloak.ExtractUserIdFromToken(w, r, handler.config.KeyCloakConfig)

	pt := &models.PrivateToken{
		UserID:       &userId,
		PrivateToken: &accessToken,
		ItemId:       &itemId,
	}

	err = users.CreateAccountToken(handler.config, pt)
	if err != nil {
		fmt.Println("Error calling user service to create account token", err)
		tools.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Long term do not want to send back to front end, just store in db and give success response
	tools.RespondNoBody(w, http.StatusCreated)
}

func (handler *Handler) CreateLinkToken(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	plaidConfig := handler.config.Plaid

	countryCodes := convertCountryCodes(strings.Split(plaidConfig.CountryCodes, ","))
	redirectURI := plaidConfig.RedirectUrl

	userId := moneymakergocloak.ExtractUserIdFromToken(w, r, handler.config.KeyCloakConfig)

	user := plaid.LinkTokenCreateRequestUser{
		ClientUserId: userId,
	}

	request := plaid.NewLinkTokenCreateRequest(
		"account-link-service-service", // Change to get this dynamically
		"en",
		countryCodes,
		user,
	)

	products := convertProducts(strings.Split(plaidConfig.Products, ","))
	request.SetProducts(products)

	if redirectURI != "" {
		request.SetRedirectUri(redirectURI)
	}

	linkTokenCreateResp, _, err := plaidConfig.Client.PlaidApi.LinkTokenCreate(ctx).LinkTokenCreateRequest(*request).Execute()
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
