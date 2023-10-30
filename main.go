package main

import (
	"hintruder/pkg/hintruder"
)

func main() {
	// cmd.Execute()
	request := hintruder.ParserRequest("test.txt")
	request.Https = true
	// print(request.Do())
	request.Do()
}
