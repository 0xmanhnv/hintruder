package hintruder

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// REF: https://github.com/xscorp/Burpee/blob/master/burpee.py
func ParserRequest(_filename string) Request {
	headerCollectionDone := false
	var request = Request{}
	var headers = make(map[string]string)

	f, err := os.Open(_filename)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan() //https://stackoverflow.com/questions/54614588/how-can-i-skip-the-first-line-of-a-file-in-go
	resource := scanner.Text()
	request.Method = strings.TrimSpace(strings.Split(resource, " ")[0])
	request.Path = strings.TrimSpace(strings.Split(resource, " ")[1])

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

				headers[k] = v
			}
		} else {
			// Collection data
			request.Data += line
		}
	}
	request.Headers = headers

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return request
}
