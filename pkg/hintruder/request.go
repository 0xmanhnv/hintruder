package hintruder

import (
	"bytes"
	"fmt"
	"net/http"
)

type Request struct {
	Https   bool
	Method  string
	Host    string
	Path    string
	Headers map[string]string
	Data    string
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
func (r Request) Do() {
	var url string
	if r.Https {
		url = fmt.Sprintf("https://%s%s", r.Host, r.Path)
	} else {
		url = fmt.Sprintf("http://%s%s", r.Host, r.Path)
	}

	req, err := http.NewRequest(r.Method, url, bytes.NewBuffer([]byte(r.Data)))
	if err != nil {
		panic(err)
	}

	for k, v := range r.Headers {
		req.Header.Add(k, v)
	}

	// Create Client
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println(res.Body)
}
