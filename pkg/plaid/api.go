package plaid

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	httptoolbox "github.com/jaydamon/http-toolbox"
	"github.com/jaydamon/moneymakergocloak"
	"github.com/plaid/plaid-go/plaid"
)

type PublicToken struct {
	PublicToken string `json:"publicToken"`
}

type PrivateToken struct {
	PrivateToken *string `json:"privateToken"`
	ItemId       *string `json:"itemId"`
}

type LinkToken struct {
	LinkToken *string `json:"linkToken"`
}

func (config *PlaidConfig) CreatePrivateAccessToken(w http.ResponseWriter, r *http.Request) {

	var tools httptoolbox.Tools
	var publicToken PublicToken

	log.Print("Recieved request ", r.Body)

	err := json.NewDecoder(r.Body).Decode(&publicToken)
	if err != nil {
		tools.ErrorJSON(w, err, 500)
		return
	}

	ctx := context.Background()

	client := config.Client
	exchangePublicTokenResp, _, err := client.PlaidApi.ItemPublicTokenExchange(ctx).ItemPublicTokenExchangeRequest(
		*plaid.NewItemPublicTokenExchangeRequest(publicToken.PublicToken),
	).Execute()
	if err != nil {
		tools.ErrorJSON(w, err, 500)
		return
	}

	accessToken := exchangePublicTokenResp.GetAccessToken() // This needs to be saved in the database
	itemId := exchangePublicTokenResp.GetItemId()

	pt := &PrivateToken{
		PrivateToken: &accessToken,
		ItemId:       &itemId,
	}

	// Long term do not want to send back to front end, just store in db and give success response
	_ = tools.WriteJSON(w, http.StatusOK, pt)
}

func (config *PlaidConfig) CreateLinkToken(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var tools httptoolbox.Tools

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
		tools.ErrorJSON(w, err, 500)
		return
	}

	linkToken := linkTokenCreateResp.GetLinkToken()
	tools.WriteJSON(w, 200, LinkToken{LinkToken: &linkToken})
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
