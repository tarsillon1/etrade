package etrade

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/tarsillon1/oauth1"
)

type httpClientSource = func(ctx context.Context) (*http.Client, error)

func do[Output any](ctx context.Context, httpClientSource httpClientSource, method string, url, nestedField string, input any) (Output, error) {
	var empty Output
	client, err := httpClientSource(ctx)
	if err != nil {
		return empty, err
	}

	var body io.Reader
	if input != nil {
		b, err := json.Marshal(input)
		if err != nil {
			return empty, err
		}
		body = bytes.NewReader(b)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return empty, err
	}
	req = req.WithContext(ctx)

	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return empty, err
	}
	defer res.Body.Close()

	b, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode != 200 {
		return empty, fmt.Errorf("request failed with status code %d: %s", res.StatusCode, string(b))
	}

	var nested map[string]Output
	err = json.Unmarshal(b, &nested)
	if err != nil {
		return empty, fmt.Errorf("failed to parse response body: %e \n %s", err, string(b))
	}

	return nested[nestedField], err
}

const (
	ProductionAPIUrl = "https://api.etrade.com"
	SandboxAPIUrl    = "https://apisb.etrade.com"
)

type Client struct {
	*Config
	Account *accountClient
	Order   *orderClient
}

func (c *Client) getHttpClient(ctx context.Context) (*http.Client, error) {
	token, err := c.Oauth1TokenSource.Token()
	if err != nil {
		return nil, err
	}
	if token == nil {
		return nil, errors.New("missing access token")
	}
	client := c.OAuth1Config.Client(ctx, token)
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			MaxVersion: tls.VersionTLS12,
		},
	}
	return client, nil
}

// Config for etrade client.
type Config struct {
	APIUrl            string
	OAuth1Config      oauth1.Config
	Oauth1TokenSource oauth1.TokenSource
}

// New creates an etrade client.
func New(config *Config) *Client {
	client := &Client{
		Config: config,
	}
	client.Account = &accountClient{apiUrl: config.APIUrl, httpClientSource: client.getHttpClient}
	client.Order = &orderClient{apiUrl: config.APIUrl, httpClientSource: client.getHttpClient}
	return client
}
