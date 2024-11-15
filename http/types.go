package http

import (
	"time"
)

type RequestConfig struct {
	Timeout  time.Duration
	MaxConns int
	Headers  map[string]string
}

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
)

type HttpClientInterface interface {
	Execute(
		method Method,
		url string,
		headers map[string]string,
		params map[string]string,
		body []byte,
		callback func() any,
	) (any, error)
}

type DefaultHttpFactoryInterface interface {
	Create(config *RequestConfig) *HttpClientInterface
}
