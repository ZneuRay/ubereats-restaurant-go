package ubereats

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Token struct {
	client      *Client
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   uint32 `json:"expires_in"`
	Scope       string `json:"scope"`
}

func (c *Client) GetToken(loginURL string, clientId string, clientSecret string, scope string) (token *Token, err error) {

	if loginURL == "" {
		return token, fmt.Errorf("Invalid login url")
	}
	if clientId == "" {
		return token, fmt.Errorf("Invalid client id")
	}
	if clientSecret == "" {
		return token, fmt.Errorf("Invalid client secret")
	}
	if scope == "" {
		return token, fmt.Errorf("Invalid scope")
	}

	params := url.Values{}
	params.Set("client_id", clientId)
	params.Set("client_secret", clientSecret)
	params.Set("grant_type", "client_credentials")
	params.Set("scope", scope)

	req, err := http.NewRequest("POST", loginURL, strings.NewReader(params.Encode()))
	if err != nil {
		return token, fmt.Errorf("Cannot create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	err = c.do(req, &token)
	if err != nil {
		return token, err
	}

	if token != nil {
		token.client = c
	}

	return
}
