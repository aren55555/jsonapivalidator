package test

import (
	"testing"

	"github.com/aren55555/jsonapivalidator"
)

func TestValidate_validErrors(t *testing.T) {
	data := []byte(`{
		"errors": [
		  {
		    "status": "422",
		    "source": { "pointer": "/data/attributes/first-name" },
		    "title":  "Invalid Attribute",
		    "detail": "First name must contain at least three characters."
		  }
		]
	}`)

	expectedResult(t, data, nil)
}

func TestValidate_errorsNotArray(t *testing.T) {
	data := []byte(`{
	  "errors": 32
	}`)

	expectedResult(t, data, jsonapivalidator.ErrInvalidErrorsType)
}

func TestValidate_errorNotErrorObject(t *testing.T) {
	data := []byte(`{
	  "errors": [32]
	}`)

	expectedResult(t, data, jsonapivalidator.ErrNotErrorObject)
}

func TestValidate_errorsKeys(t *testing.T) {
	data := []byte(`{
	  "errors": [{
			"aren55555": "foo"
		}]
	}`)

	expectedResult(t, data, jsonapivalidator.ErrInvalidErrorMember)
}
