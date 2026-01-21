package client

import (
	"net/http"

	"github.com/go-logr/logr"
	"github.com/go-logr/logr/funcr"
)


type RestClient struct {
	baseUrl string
	bearerToken string
	basicAuth *BasicAuth
	logger *logr.Logger
	httpClient *http.Client
}

type BasicAuth struct {
	username string
	password string
}

func NewRestClient(baseUrl string) *RestClient {
	noopLogger := funcr.New(func(prefix, args string) {}, funcr.Options{})
	return &RestClient{
		baseUrl: baseUrl,
		httpClient: http.DefaultClient,
		logger: &noopLogger, // default noop logger
	}
}

func (c *RestClient) SetBearerToken(token string) {
	c.bearerToken = token
}

func (c *RestClient) SetBasicAuth(username string, password string) {
	c.basicAuth = &BasicAuth{
		username: username,
		password: password,
	}
}

func (c *RestClient) SetLogger(logger *logr.Logger) {
	c.logger = logger
}

func (c *RestClient) SetHttpClient(client *http.Client) {
	c.httpClient = client
}