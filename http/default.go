package http

import t "net/http"

type DefaultHttpFactory struct {
}

// net/http as a default implementation
type DefaultHttpClient struct {
	client *t.Client
}

func (c *DefaultHttpClient) Execute(
	method Method,
	url string,
	headers map[string]string,
	params map[string]string,
	body []byte,
	callback func() any,
) (any, error) {
	return callback(), nil
}

func (f *DefaultHttpFactory) Create(config *RequestConfig) *DefaultHttpClient {
	return nil
}
