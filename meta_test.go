package jsonapivalidator

import "testing"

func TestValidate_meta(t *testing.T) {
	data := []byte(`{
  	"meta": {
    	"copyright": "Copyright 2015 Example Corp.",
    	"authors": [
      	"Yehuda Katz",
      	"Steve Klabnik",
      	"Dan Gebhardt",
      	"Tyler Kellen"
    	]
  	}
  }`)

	if validatePayload(t, data).HasErrors() {
		t.Fatal(testErrorNotExpected)
	}
}
