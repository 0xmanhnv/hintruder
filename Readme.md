# HOST INTRUDER



## TODO:
```go
// implement fasthttp in request
func (r Request) Do() ([]byte, error) {
	// usage
	client := &fasthttp.Client{
		ReadTimeout:         30 * time.Second,
		MaxConnsPerHost:     1024 * 3,
		MaxIdleConnDuration: -1,
		ReadBufferSize:      1024 * 8,
		TLSConfig:           &tls.Config{InsecureSkipVerify: true},
		Dial:                fasthttpproxy.FasthttpHTTPDialer("127.0.0.1:8080"),
	}

	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(r.Method)

	// req.Header.SetHost(r.Host)
	// req.URI().SetHost(r.Host)
	// req.URI().SetPath(r.Path)

	// Set request URI
	if r.Https {
		req.SetRequestURI(fmt.Sprintf("https://%s%s", r.Host, r.Path))
	} else {
		req.SetRequestURI(fmt.Sprintf("http://%s%s", r.Host, r.Path))
	}

	// req.SetTimeout(time.Duration(300) * time.Millisecond)
	// Add header
	for k, v := range r.Headers {
		req.Header.Set(k, v)
	}
	req.SetBodyRaw([]byte(r.Data))
	defer fasthttp.ReleaseRequest(req)

	fmt.Println(req.Host())

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	// Perform the request
	if err := client.Do(req, resp); err != nil {
		log.Printf("Client get failed: %s\n", err)
		return nil, err
	}

	return GetResponseBody(resp)
}

// GetResponseBody return plain response body of resp
func GetResponseBody(resp *fasthttp.Response) ([]byte, error) {
	var contentEncoding = string(resp.Header.Peek("Content-Encoding"))
	if len(contentEncoding) < 1 {
		return resp.Body(), nil
	}
	if contentEncoding == "br" {
		return resp.BodyUnbrotli()
	}
	if contentEncoding == "gzip" {
		return resp.BodyGunzip()
	}
	if contentEncoding == "deflate" {
		return resp.BodyInflate()
	}
	return nil, errors.New("unsupported response content encoding: " + string(contentEncoding))
}
```