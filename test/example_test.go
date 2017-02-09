package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/aren55555/jsonapivalidator"
)

const (
	sampleLocation = "https://raw.githubusercontent.com/aren55555/jsonapivalidator/master/test/samples/valid/default.json"
)

func TestExample(t *testing.T) {
	req, err := http.DefaultClient.Get(sampleLocation)
	if err != nil {
		t.Fatal(err)
	}
	defer req.Body.Close()

	result, err := jsonapivalidator.UnmarshalAndValidate(req.Body)
	if err != nil {
		t.Fatal(err)
	}

	if result.IsValid() {
		fmt.Println("The JSON sample was valid")
		return
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
