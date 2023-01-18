package plaid

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/factotum/moneymaker/plaid-integration/pkg/models"
	"github.com/factotum/moneymaker/plaid-integration/pkg/users"
	tools "github.com/jaydamon/http-toolbox"
	"github.com/jaydamon/moneymakergocloak"
	"github.com/plaid/plaid-go/plaid"
)

func (plaidCtx *PlaidContext) CreatePrivateAccessToken(w http.ResponseWriter, r *http.Request) {

	var publicToken models.PublicToken

	config := plaidCtx.context.Config.Plaid

	log.Print("Recieved request ", &r.Body)

	err := json.NewDecoder(r.Body).Decode(&publicToken)
	if err != nil {
		fmt.Println("Error decoading input body", err)
		tools.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ctx := context.Background()

	client := config.Client
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
	userId := moneymakergocloak.ExtractUserIdFromToken(w, r, plaidCtx.context.Config.KeyCloakConfig)

	pt := &models.PrivateToken{
		UserID:       &userId,
		PrivateToken: &accessToken,
		ItemId:       &itemId,
	}

	err = users.CreateAccountToken(plaidCtx.context.Config, pt)
	if err != nil {
		fmt.Println("Error calling user service to create account token", err)
		tools.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Long term do not want to send back to front end, just store in db and give success response
	tools.RespondNoBody(w, http.StatusCreated)
}

func (plaidCtx *PlaidContext) CreateLinkToken(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	config := plaidCtx.context.Config.Plaid

	countryCodes := convertCountryCodes(strings.Split(config.CountryCodes, ","))
	redirectURI := config.RedirectUrl

	userId := moneymakergocloak.ExtractUserIdFromToken(w, r, config.Auth)

	user := plaid.LinkTokenCreateRequestUser{
		ClientUserId: userId,
	}

	request := plaid.NewLinkTokenCreateRequest(
		"plaid-integration-service", // Change to get this dynamically
		"en",
		countryCodes,
		user,
	)

	products := convertProducts(strings.Split(config.Products, ","))
	request.SetProducts(products)

	if redirectURI != "" {
		request.SetRedirectUri(redirectURI)
	}

	linkTokenCreateResp, _, err := config.Client.PlaidApi.LinkTokenCreate(ctx).LinkTokenCreateRequest(*request).Execute()
	if err != nil {
		log.Print("Error retrieving LinkToken ", err)
		tools.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	linkToken := linkTokenCreateResp.GetLinkToken()
	tools.Respond(w, http.StatusOK, models.LinkToken{LinkToken: &linkToken})
}

func convertCountryCodes(countryCodeStrs []string) []plaid.CountryCode {
	countryCodes := []plaid.CountryCode{}

	for _, countryCodeStr := range countryCodeStrs {
		countryCodes = append(countryCodes, plaid.CountryCode(countryCodeStr))
	}

	return countryCodes
}

func convertProducts(productStrs []string) []plaid.Products {
	products := []plaid.Products{}

	for _, productStr := range productStrs {
		products = append(products, plaid.Products(productStr))
	}

	return products
}
