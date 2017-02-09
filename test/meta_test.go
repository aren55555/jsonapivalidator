package test

import (
	"testing"

	"github.com/aren55555/jsonapivalidator"
)

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

	expectedResult(t, data, noError, noWarning)
}

func TestValidate_invalidMeta(t *testing.T) {
	data := []byte(`{
  	"meta": 21
  }`)

	expectedResult(t, data, jsonapivalidator.ErrNotMetaObject, noWarning)
}
