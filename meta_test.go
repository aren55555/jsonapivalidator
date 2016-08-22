package jsonapivalidator

import "testing"

func TestValidate_validMeta(t *testing.T) {
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

func TestValidate_invalidMeta(t *testing.T) {
	data := []byte(`{
  	"meta": 21
  }`)

	if expecting, r := ErrNotMetaObject, validatePayload(t, data); !r.HasError(expecting) {
		t.Fatalf(testErrorExpected, expecting, r.Errors())
	}
}
