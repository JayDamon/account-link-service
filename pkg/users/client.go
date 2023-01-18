package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/factotum/moneymaker/plaid-integration/pkg/config"
	"github.com/factotum/moneymaker/plaid-integration/pkg/models"
)

func CreateAccountToken(config *config.Config, token *models.PrivateToken) error {

	uri := fmt.Sprintf("%s/api/v1/users/%s/account-tokens", config.UserServiceUrl, *token.UserID)

	body, _ := json.Marshal(token)

	resp, err := http.Post(uri, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {

		msg := fmt.Sprintf("Unexpected status code found on Token Creation.Code: \n%sBody: \n%s", resp.Status, resp.Body)

		return fmt.Errorf(msg)
	}

	return nil
}
