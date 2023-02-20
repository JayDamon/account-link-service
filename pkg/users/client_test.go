package users

import (
	"fmt"
	"github.com/factotum/moneymaker/account-link-service/pkg/config"
	"github.com/factotum/moneymaker/account-link-service/pkg/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_CreateAccountToken_ExpectStatusCrated(t *testing.T) {

	userId := "testId"
	testToken := createPrivateAccountToken(&userId)

	urlPath := fmt.Sprintf("/v1/users/%s/account-tokens", userId)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path != urlPath {
			t.Errorf("Expected to request '%s', got: %s", urlPath, r.URL.Path)
		}
		w.WriteHeader(http.StatusCreated)
	}))
	defer server.Close()

	testConfig := config.Config{
		UserServiceUrl: server.URL,
	}

	err := CreateAccountToken(&testConfig, testToken)
	assert.Nil(t, err, "Unexpected err %s", err)
}

func TestCreateAccountToken_NoServiceUrl(t *testing.T) {

	userId := "testId"
	testToken := createPrivateAccountToken(&userId)

	urlPath := fmt.Sprintf("/api/v1/users/%s/account-tokens", userId)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != urlPath {
			t.Errorf("Expected to request '%s', got: %s", urlPath, r.URL.Path)
		}
		w.WriteHeader(http.StatusCreated)
	}))
	defer server.Close()

	testConfig := config.Config{}

	err := CreateAccountToken(&testConfig, testToken)
	assert.NotNil(t, err, "Expected err, but was nil")
	expectedErr := "expected valid UserServiceUrl, found nil or empty string"
	if err.Error() != expectedErr {
		t.Errorf("Expected err '%s', found '%s'", expectedErr, err)
	}
}

func TestCreateAccountToken_NilUserIdProvided(t *testing.T) {
	testToken := createPrivateAccountToken(nil)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	}))
	defer server.Close()

	testConfig := config.Config{
		UserServiceUrl: server.URL,
	}

	err := CreateAccountToken(&testConfig, testToken)
	assert.NotNil(t, err, "Expected err, but was nil")

	expectedErr := "expected valid UserId, found nil or empty string"
	if err.Error() != expectedErr {
		t.Errorf("Expected err '%s', found '%s'", expectedErr, err)
	}
}

func TestCreateAccountToken_EmptyStringUserIdProvided(t *testing.T) {
	userId := ""
	testToken := createPrivateAccountToken(&userId)

	urlPath := fmt.Sprintf("/api/v1/users/%s/account-tokens", userId)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != urlPath {
			t.Errorf("Expected to request '%s', got: %s", urlPath, r.URL.Path)
		}
		w.WriteHeader(http.StatusCreated)
	}))
	defer server.Close()

	testConfig := config.Config{
		UserServiceUrl: server.URL,
	}

	err := CreateAccountToken(&testConfig, testToken)
	assert.NotNil(t, err, "Expected err, but was nil")

	expectedErr := "expected valid UserId, found nil or empty string"
	if err.Error() != expectedErr {
		t.Errorf("Expected err '%s', found '%s'", expectedErr, err)
	}
}

func TestCreateAccountToken_NilPrivateTokenProvided(t *testing.T) {
	userId := "testToken"
	testToken := createPrivateAccountToken(&userId)
	testToken.PrivateToken = nil

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	}))
	defer server.Close()

	testConfig := config.Config{
		UserServiceUrl: server.URL,
	}

	err := CreateAccountToken(&testConfig, testToken)
	assert.NotNil(t, err, "Expected err, but was nil")

	expectedErr := "expected valid Private Token, found nil or empty string"
	if err.Error() != expectedErr {
		t.Errorf("Expected err '%s', found '%s'", expectedErr, err)
	}
}

func TestCreateAccountToken_EmptyStringPrivateTokenProvided(t *testing.T) {
	userId := "testToken"
	testToken := createPrivateAccountToken(&userId)
	privateToken := ""
	testToken.PrivateToken = &privateToken

	urlPath := fmt.Sprintf("/api/v1/users/%s/account-tokens", userId)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != urlPath {
			t.Errorf("Expected to request '%s', got: %s", urlPath, r.URL.Path)
		}
		w.WriteHeader(http.StatusCreated)
	}))
	defer server.Close()

	testConfig := config.Config{
		UserServiceUrl: server.URL,
	}

	err := CreateAccountToken(&testConfig, testToken)
	assert.NotNil(t, err, "Expected err, but was nil")

	expectedErr := "expected valid Private Token, found nil or empty string"
	if err.Error() != expectedErr {
		t.Errorf("Expected err '%s', found '%s'", expectedErr, err)
	}
}

func TestCreateAccountToken_NilItemIdProvided(t *testing.T) {
	userId := "testToken"
	testToken := createPrivateAccountToken(&userId)
	testToken.ItemId = nil

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	}))
	defer server.Close()

	testConfig := config.Config{
		UserServiceUrl: server.URL,
	}

	err := CreateAccountToken(&testConfig, testToken)
	assert.NotNil(t, err, "Expected err, but was nil")

	expectedErr := "expected valid Item Id, found nil or empty string"
	if err.Error() != expectedErr {
		t.Errorf("Expected err '%s', found '%s'", expectedErr, err)
	}
}

func TestCreateAccountToken_EmptyStringItemIdProvided(t *testing.T) {
	userId := "testToken"
	testToken := createPrivateAccountToken(&userId)
	itemId := ""
	testToken.ItemId = &itemId

	urlPath := fmt.Sprintf("/api/v1/users/%s/account-tokens", userId)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != urlPath {
			t.Errorf("Expected to request '%s', got: %s", urlPath, r.URL.Path)
		}
		w.WriteHeader(http.StatusCreated)
	}))
	defer server.Close()

	testConfig := config.Config{
		UserServiceUrl: server.URL,
	}

	err := CreateAccountToken(&testConfig, testToken)
	assert.NotNil(t, err, "Expected err, but was nil")

	expectedErr := "expected valid Item Id, found nil or empty string"
	if err.Error() != expectedErr {
		t.Errorf("Expected err '%s', found '%s'", expectedErr, err)
	}
}

func TestCreateAccountToken_StatusCreatedNotReturned(t *testing.T) {
	userId := "testId"
	testToken := createPrivateAccountToken(&userId)

	urlPath := fmt.Sprintf("/v1/users/%s/account-tokens", userId)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != urlPath {
			t.Errorf("Expected to request '%s', got: %s", urlPath, r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	testConfig := config.Config{
		UserServiceUrl: server.URL,
	}

	err := CreateAccountToken(&testConfig, testToken)
	assert.NotNil(t, err, "Expected err, but was nil")

	expectedErr := "unexpected status code found on Token Creation.\nCode: 200 OK\nBody: {}"
	if err.Error() != expectedErr {
		t.Errorf("Expected err '%s', found '%s'", expectedErr, err)
	}
}

func createPrivateAccountToken(userId *string) *models.PrivateToken {
	testToken := "testToken"
	testItemId := "testItemId"
	return &models.PrivateToken{
		UserID:       userId,
		PrivateToken: &testToken,
		ItemId:       &testItemId,
	}
}
