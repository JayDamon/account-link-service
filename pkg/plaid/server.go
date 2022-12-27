package plaid

// func createLinkToken(c *gin.Context) {
//   ctx := context.Background()
//   // Get the client_user_id by searching for the current user
//   user, _ := usermodels.Find(...)
//   clientUserId := user.ID.String()
//   // Create a link_token for the given user
//   request := plaid.NewLinkTokenCreateRequest("Plaid Test App", "en", []plaid.CountryCode{plaid.COUNTRYCODE_US}, *plaid.NewLinkTokenCreateRequestUser(clientUserId))
//   request.SetWebhook("https://webhook.sample.com")
//   request.SetRedirectUri("https://domainname.com/oauth-page.html")
//   request.SetProducts([]plaid.Products{plaid.PRODUCTS_AUTH})
//     resp, _, err := testClient.PlaidApi.LinkTokenCreate(ctx).LinkTokenCreateRequest(*request).Execute()
//   // Send the data to the client
//   c.JSON(http.StatusOK, gin.H{
//     "link_token": resp.GetLinkToken(),
//   })
// }
