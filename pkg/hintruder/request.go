package hintruder

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"reflect"
)

type Request struct {
	Protocol    string `default:"http"`
	Method      string `default:"GET"`
	Host        string
	Path        string `default:"/"`
	Headers     map[string]string
	Data        string
	TlsVerify   bool
	ProxyUrl    *url.URL
	ProxyEnable bool
	http.Request
}

func setDefaults(p *Request) {
	// Iterate over the fields of the Person struct using reflection
	// and set the default value for each field if the field is not provided
	// by the caller of the constructor function.
	for i := 0; i < reflect.TypeOf(*p).NumField(); i++ {
		field := reflect.TypeOf(*p).Field(i)
		if value, ok := field.Tag.Lookup("default"); ok {
			switch field.Type.Kind() {
			case reflect.String:
				if p.Protocol == "" {
					p.Protocol = value
				}
				if p.Method == "" {
					p.Method = value
				}
			case reflect.Bool:
				if p.ProxyEnable && p.ProxyUrl.String() == "" {
					proxyUrl, _ := url.Parse("http://127.0.0.1:8081")
					p.ProxyUrl = proxyUrl
				}
			}

		}
	}
}

func (r *Request) AddHeader(headers map[string]string) *Request {
	r.Headers = headers
	r.Host = headers["Host"]
	return r
}

// https://www.kirandev.com/http-post-golang
// https://golangnote.com/request/sending-post-request-in-golang-with-header
// https://thedevelopercafe.com/articles/make-post-request-in-go-d9756284d70b
// https://www.golangprograms.com/how-do-you-set-headers-in-an-http-request-with-an-http-client-in-go.html
func (r Request) Do() (*http.Response, error) {
	urlRequest := fmt.Sprintf("%s://%s%s", r.Protocol, r.Host, r.Path)
	req, err := http.NewRequest(r.Method, urlRequest, bytes.NewBuffer([]byte(r.Data)))
	if err != nil {
		return &http.Response{}, err
	}

	// Add headers
	for k, v := range r.Headers {
		req.Header.Add(k, v)
	}

	// Create Client
	client := &http.Client{
		Transport: &http.Transport{
			Proxy:           http.ProxyURL(r.ProxyUrl),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: r.TlsVerify},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return &http.Response{}, err
	}
	// Close response body
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	return resp, nil
}
