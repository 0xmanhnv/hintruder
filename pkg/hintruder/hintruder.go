package hintruder

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"net/url"
)

type Hintruder struct {
	Protocol  string
	ProxyUrl  string
	TlsVerify bool
}

func (h *Hintruder) Run(ctx context.Context, _fileRequest string) *http.Response {
	req, err := ParserRequest(ctx, _fileRequest)
	if err != nil {
		log.Fatal("Parse file failed!!!")
	}

	req.URL.Scheme = h.Protocol
	if h.Protocol == "" {
		req.URL.Scheme = "http"
	}

	// Create Client
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: h.TlsVerify},
	}

	// Add proxy
	if h.ProxyUrl != "" {
		proxyUrl, _ := url.Parse(h.ProxyUrl)
		transport.Proxy = http.ProxyURL(proxyUrl)
	}

	client := &http.Client{
		Transport: transport,
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do request failed!!!", err)
	}
	// Close response body
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	return resp
}
