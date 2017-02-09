package main

import (
	"fmt"
	"net/http"

	"github.com/aren55555/jsonapivalidator"
)

func main() {
	// First file is valid.
	req, err := http.DefaultClient.Get("https://raw.githubusercontent.com/aren55555/jsonapivalidator/master/test/samples/valid/default.json")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	result, err := jsonapivalidator.UnmarshalAndValidate(req.Body)
	if err != nil {
		panic(err)
	}

	if result.IsValid() {
		fmt.Println("The JSON sample was valid!")
	}

	// Second file is invalid
	req, err = http.DefaultClient.Get("https://raw.githubusercontent.com/aren55555/jsonapivalidator/master/test/samples/invalid/default.json")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	result, err = jsonapivalidator.UnmarshalAndValidate(req.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Errors:")
	for i, err := range result.Errors() {
		fmt.Println("\t", i, err)
	}

	fmt.Println("Warnings:")
	for i, err := range result.Warnings() {
		fmt.Println("\t", i, err)
	}
}
