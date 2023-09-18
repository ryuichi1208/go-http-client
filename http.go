package gohttpclient

import "net/http"

type HttpClient struct {
	client http.Client
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		client: http.Client{
			Transport: &http.Transport{
				Proxy:                  nil,
				DialContext:            nil,
				MaxIdleConnsPerHost:    100,
				MaxIdleConns:           100,
				IdleConnTimeout:        0,
				TLSHandshakeTimeout:    0,
				ExpectContinueTimeout:  0,
				ResponseHeaderTimeout:  0,
				DisableKeepAlives:      false,
				DisableCompression:     false,
				MaxResponseHeaderBytes: 0,
				WriteBufferSize:        0,
				ReadBufferSize:         0,
				ForceAttemptHTTP2:      false,
				TLSClientConfig:        nil,
				TLSNextProto:           nil,
				ProxyConnectHeader:     nil,
				MaxConnsPerHost:        0,
			},
			CheckRedirect: nil,
			Timeout:       0,
		},
	}
}

func (h *HttpClient) Get(url string) (resp *http.Response, err error) {
	return h.client.Get(url)
}

func (h *HttpClient) Head(url string) (resp *http.Response, err error) {
	return h.client.Head(url)
}
