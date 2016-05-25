package wikidata

import (
	"net/http"
)

// Config for the WikiData clients.
type Config struct {
	HTTPClient *http.Client
}

// NewConfig with the given access token.
func NewConfig() *Config {
	return &Config{
		HTTPClient: http.DefaultClient,
	}
}
