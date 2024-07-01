package etrade

import (
	"fmt"

	"github.com/tarsillon1/oauth1"
)

const (
	apiUrl          = "https://api.etrade.com"
	authorizeUrl    = "https://us.etrade.com/e/t/etws/authorize?oauth_consumer_key=%s"
	accessTokenUrl  = apiUrl + "/oauth/access_token"
	requestTokenUrl = apiUrl + "/oauth/request_token"
)

// NewOAuth1Config creates an OAuth configuration with pre-configured endpoints for ETrade OAuth provider.
func NewOAuth1Config(consumerKey, consumerSecret, callbackUrl string) oauth1.Config {
	return oauth1.Config{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		CallbackURL:    callbackUrl,
		Endpoint: oauth1.Endpoint{
			AccessTokenURL:  accessTokenUrl,
			RequestTokenURL: requestTokenUrl,
			AuthorizeURL:    fmt.Sprintf(authorizeUrl, consumerKey),
		},
		HTTPMethod: "GET",
	}
}
