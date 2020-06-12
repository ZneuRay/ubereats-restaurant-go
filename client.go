package ubereats

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	client  *http.Client
	Token   string
	BaseURL string
}

func NewClient(baseURL string, Token string) *Client {
	c := &Client{
		client:  http.DefaultClient,
		Token:   Token,
		BaseURL: baseURL,
	}

	return c
}

func (c *Client) SetBaseURL(baseURL string) {
	c.BaseURL = baseURL
}

func (c *Client) Get(path string, target interface{}) error {
	url := fmt.Sprintf("%s/%s", c.BaseURL, path)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("Cannot create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))

	return c.do(req, target)
}

func (c *Client) Post(path string, data, response interface{}) error {
	url := fmt.Sprintf("%s/%s", c.BaseURL, path)

	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("Cannot marshal body: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("Cannot create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))

	return c.do(req, response)
}

func (c *Client) do(req *http.Request, response interface{}) error {
	if req == nil {
		return errors.New("nil request")
	}

	url := req.URL.RequestURI()
	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("HTTP request failure on %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return makeHTTPClientError(url, resp)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("HTTP Read error on response for %s: %w", url, err)
	}

	err = json.Unmarshal(b, response)
	if err != nil {
		return fmt.Errorf("JSON decode failed on %s:\n%s\nerror: %w", url, string(b), err)
	}
	return nil
}
