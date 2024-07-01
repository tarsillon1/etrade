package etrade

import (
	"fmt"

	"github.com/tarsillon1/oauth1"
)

const (
	apiUrl          = "https://api.etrade.com"
	authorizeUrl    = "https://us.etrade.com/e/t/etws/authorize?key=%s"
	accessTokenUrl  = apiUrl + "/oauth/access_token"
	requestTokenUrl = apiUrl + "/oauth/request_token"
)

// NewOAuth1Config creates an OAuth configuration with pre-configured endpoints for ETrade OAuth provider.
func NewOAuth1Config(consumerKey, consumerSecret string) oauth1.Config {
	return oauth1.Config{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		CallbackURL:    "oob",
		Endpoint: oauth1.Endpoint{
			AccessTokenURL:  accessTokenUrl,
			RequestTokenURL: requestTokenUrl,
			AuthorizeURL:    fmt.Sprintf(authorizeUrl, consumerKey),
		},
		HTTPMethod:                   "GET",
		AuthorizationOAuthTokenParam: "token",
	}
}
