package plaidlink

import (
	"context"
	"github.com/plaid/plaid-go/plaid"
	"net/http"
)

type PlaidAccountApi struct {
	plaidApi plaid.PlaidApiService
}

type ApiService interface {
	ItemPublicTokenExchange(ctx context.Context, tokenExchangeRequest *plaid.ItemPublicTokenExchangeRequest) (plaid.ItemPublicTokenExchangeResponse, *http.Response, error)
	RequestLinkToken(ctx context.Context, linkTokenRequest *plaid.LinkTokenCreateRequest) (plaid.LinkTokenCreateResponse, *http.Response, error)
}

func (api *PlaidAccountApi) ItemPublicTokenExchange(
	ctx context.Context,
	tokenExchangeRequest *plaid.ItemPublicTokenExchangeRequest,
) (plaid.ItemPublicTokenExchangeResponse, *http.Response, error) {
	return api.plaidApi.ItemPublicTokenExchange(ctx).ItemPublicTokenExchangeRequest(*tokenExchangeRequest).Execute()
}

func (api *PlaidAccountApi) RequestLinkToken(
	ctx context.Context,
	linkTokenRequest *plaid.LinkTokenCreateRequest,
) (plaid.LinkTokenCreateResponse, *http.Response, error) {
	return api.plaidApi.LinkTokenCreate(ctx).LinkTokenCreateRequest(*linkTokenRequest).Execute()
}

func NewApiService(config *plaid.Configuration) ApiService {
	return &PlaidAccountApi{
		plaidApi: *plaid.NewAPIClient(config).PlaidApi,
	}
}
