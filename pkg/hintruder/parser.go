package hintruder

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// REF: https://github.com/xscorp/Burpee/blob/master/burpee.py
func ParserRequest(_filename string) {
	headerCollectionDone := false
	// TODO: Read File
	f, err := os.Open(_filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Text()

	for scanner.Scan() {
		line := scanner.Text()

		if !headerCollectionDone {
			if line == "" {
				headerCollectionDone = true
				fmt.Println("End header")
			} else {
				fmt.Println("Collection header")
				fmt.Println(strings.Split(line, ":"))
				Headers[strings.Split(line, ":")[0]] = strings.Split(line, ":")[1]
			}
		}
	}

	fmt.Println(Headers)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// TODO: Parse to header
	// TODO: Parse to body
}
