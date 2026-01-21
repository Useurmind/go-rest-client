package client

import (
	"github.com/google/go-querystring/query"
	"net/url"
)

func EncodeUrlValues(obj interface{}) (url.Values, error) {
	return query.Values(obj)
}