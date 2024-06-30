package etrade

import "github.com/dghubble/oauth1"

type client struct {
	config *Config
}

// Config for etrade client.
type Config struct {
	OAuth1Config      oauth1.Config
	Oauth1TokenSource oauth1.TokenSource
}

// New creates an etrade client.
func New(config *Config) *client {
	client := &client{
		config: config,
	}
	return client
}
