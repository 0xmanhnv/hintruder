package hintruder

import (
	"bufio"
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// REF: https://github.com/xscorp/Burpee/blob/master/burpee.py
func ParserRequest(ctx context.Context, _filename string) (*http.Request, error) {
	headerCollectionDone := false
	// request := http.Request{}
	// http.NewRequest()
	//setDefaults(&request) // Set default value request
	req, err := http.NewRequestWithContext(ctx, "", "", nil)
	if err != nil {
		return nil, err
	}
	var data string

	f, err := os.Open(_filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan() //https://stackoverflow.com/questions/54614588/how-can-i-skip-the-first-line-of-a-file-in-go
	resource := scanner.Text()
	req.Method = strings.TrimSpace(strings.Split(resource, " ")[0])
	req.URL.Path = strings.TrimSpace(strings.Split(resource, " ")[1])
	req.Proto = strings.TrimSpace(strings.Split(resource, " ")[2])

	for scanner.Scan() {
		line := scanner.Text() // Passed first line
		// Collection headers
		if !headerCollectionDone {
			if line == "" {
				headerCollectionDone = true
				continue
			} else {
				k := strings.TrimSpace(line[:strings.Index(line, ":")])
				v := strings.TrimSpace(line[strings.Index(line, ":")+1:])
				req.Header.Set(k, v)

				if k == "Host" {
					req.URL.Host = v
				}
			}
		} else {
			// Collection data
			data += line
		}
	}

	// Add data to request
	stringReader := strings.NewReader(data)
	req.Body = io.NopCloser(stringReader)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return req, nil
}
