package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/factotum/moneymaker/account-link-service/pkg/config"
	"github.com/factotum/moneymaker/account-link-service/pkg/models"
)

func CreateAccountToken(config *config.Config, token *models.PrivateToken, bearerToken string) error {

	err2 := validateConfigAndToken(config, token)
	if err2 != nil {
		return err2
	}

	url := fmt.Sprintf("%s/v1/users/%s/account-tokens", config.UserServiceUrl, *token.UserID)

	body, _ := json.Marshal(token)

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", bearerToken)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {

		msg := fmt.Sprintf("unexpected status code found on Token Creation.\nCode: %s\nBody: %s", resp.Status, resp.Body)

		return fmt.Errorf(msg)
	}

	return nil
}

func validateConfigAndToken(config *config.Config, token *models.PrivateToken) error {
	if config.UserServiceUrl == "" {
		return fmt.Errorf("expected valid UserServiceUrl, found nil or empty string")
	}

	if token.UserID == nil || *token.UserID == "" {
		return fmt.Errorf("expected valid UserId, found nil or empty string")
	}

	if token.PrivateToken == nil || *token.PrivateToken == "" {
		return fmt.Errorf("expected valid Private Token, found nil or empty string")
	}

	if token.ItemId == nil || *token.ItemId == "" {
		return fmt.Errorf("expected valid Item Id, found nil or empty string")
	}
	return nil
}
