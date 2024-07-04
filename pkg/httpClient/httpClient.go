package httpClient

import (
	"net/http"
	"net/url"
)

// CreateHTTPClient cria um cliente HTTP configurado para usar um proxy com autenticação básica.
func CreateHTTPClient(proxyURL *url.URL, proxyUsername, proxyPassword string) *http.Client {
	proxyURL.User = url.UserPassword(proxyUsername, proxyPassword)
	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}
}
