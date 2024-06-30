package etrade

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/dghubble/oauth1"
)

type httpClientSource = func(ctx context.Context) (*http.Client, error)

func do[Output any](ctx context.Context, httpClientSource httpClientSource, method string, url string, input any) (Output, error) {
	var target Output
	client, err := httpClientSource(ctx)
	if err != nil {
		return target, err
	}

	var body io.Reader
	if input != nil {
		b, err := json.Marshal(input)
		if err != nil {
			return target, err
		}
		body = bytes.NewReader(b)

	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return target, err
	}
	req = req.WithContext(ctx)

	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	res, err := client.Do(req)
	if err != nil {
		return target, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(target)
	if err != nil {
		return target, err
	}

	return target, err
}

const (
	ProductionAPIUrl = "https://api.etrade.com"
	SandboxAPIUrl    = "https://apisb.etrade.com"
)

type client struct {
	*Config
	Account *accountClient
	Order   *orderClient
}

func (c *client) getHttpClient(ctx context.Context) (*http.Client, error) {
	token, err := c.Oauth1TokenSource.Token()
	if err != nil {
		return nil, err
	}
	if token == nil {
		return nil, errors.New("missing access token")
	}
	return c.OAuth1Config.Client(ctx, token), nil
}

// Config for etrade client.
type Config struct {
	APIUrl            string
	OAuth1Config      oauth1.Config
	Oauth1TokenSource oauth1.TokenSource
}

// New creates an etrade client.
func New(config *Config) *client {
	client := &client{
		Config: config,
	}
	client.Account = &accountClient{apiUrl: config.APIUrl, httpClientSource: client.getHttpClient}
	client.Order = &orderClient{apiUrl: config.APIUrl, httpClientSource: client.getHttpClient}
	return client
}
