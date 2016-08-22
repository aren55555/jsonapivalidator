package test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/aren55555/jsonapivalidator"
)

func TestUnmarshalAndValidate_allValid(t *testing.T) {
	validDir := "./samples/valid"
	files, err := ioutil.ReadDir(validDir)
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range files {
		if f.IsDir() {
			continue // skip dirs
		}

		data, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", validDir, f.Name()))
		if err != nil {
			t.Fatal(err)
		}

		r, err := jsonapivalidator.UnmarshalAndValidate(data)
		if err != nil {
			t.Fatal(err)
		}

		if len(r.Errors()) > 0 {
			t.Fatal("Was not expecting any errors")
		}
	}
}
