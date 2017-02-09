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

	expectedResult(t, data, noError, noWarning)
}

func TestValidate_errorsNotArray(t *testing.T) {
	data := []byte(`{
	  "errors": 32
	}`)

	expectedResult(t, data, jsonapivalidator.ErrInvalidErrorsType, noWarning)
}

func TestValidate_errorNotErrorObject(t *testing.T) {
	data := []byte(`{
	  "errors": [32]
	}`)

	expectedResult(t, data, jsonapivalidator.ErrNotErrorObject, noWarning)
}

func TestValidate_errorsKeys(t *testing.T) {
	data := []byte(`{
	  "errors": [{
			"aren55555": "foo"
		}]
	}`)

	expectedResult(t, data, jsonapivalidator.ErrInvalidErrorMember, noWarning)
}
