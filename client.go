package beem

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
)

// Version of the api of the package
//
// Changing this value indicates introduction of breaking change
const Version = 1

// Init configuration options to create a `beem.New`
type Options struct {
	ApiKey    string
	SecretKey string
	ApiUrl    string
}

// Instance of the beem application as
// used through out the package
type App struct {
	url     string
	options Options
}

const (
	EnvVarBeemAfricaV1AppUrl  = "BEEM_AFRICA_SMS_V1_API_URL"
	EnvVarBeemAfricaApiKey    = "BEEM_AFRICA_SMS_API_KEY"
	EnvVarBeemAfricaSecretKey = "BEEM_AFRICA_SMS_SECRET_KEY"
)

// Known Beem API URL referenced in the v0.1.0 of the documentation
// Reference link: https://docs.beem.africa/#api-_
const defaultBeemApiUrl = "https://apisms.beem.africa"

// Loads the `beem.Options` from environment where the code runs
func DefaultEnvOptions() Options {
	apiUrl := os.Getenv(EnvVarBeemAfricaV1AppUrl)

	if apiUrl == "" {
		// set known default url
		apiUrl = defaultBeemApiUrl
	}

	return Options{
		ApiKey:    os.Getenv(EnvVarBeemAfricaApiKey),
		SecretKey: os.Getenv(EnvVarBeemAfricaSecretKey),
		ApiUrl:    apiUrl,
	}
}

// Get's the Api Key associated
// to the beem instance
func (b *App) ApiKey() string {
	return b.options.ApiKey
}

// Get's the Secret attached to the application
func (b *App) SecretKey() string {
	return b.options.SecretKey
}

// Get's the API URL expected of the instance
// of the application
func (b *App) ApiUrl() string {
	return b.options.ApiUrl
}

func New(opts Options) (*App, error) {
	// setup defaults
	url := opts.ApiUrl
	if url == "" {
		url = defaultBeemApiUrl
	}

	if opts.ApiKey == "" || opts.SecretKey == "" {
		return nil, errors.New("missing configuration beem configuration for `BeemApiKey` and `BeemSecretKey`. include the options using `beem.Options` or use the `beem.EnvOptions()` helper ")
	}

	return &App{
		url:     url,
		options: Options{},
	}, nil
}

// Constructs the authorization token to be attached as
// `Authorization: Basic <TOKEN>`
func (b *App) GetRequestAuthToken() string {
	return base64.StdEncoding.EncodeToString(
		[]byte(fmt.Sprintf("%s:%s", b.ApiKey(), b.SecretKey())),
	)
}
