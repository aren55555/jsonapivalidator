package jsonapivalidator

import "testing"

func TestValidate_errorsNotArray(t *testing.T) {
	data := []byte(`{
	  "errors": 32
	}`)

	if expecting, r := ErrInvalidErrorsType, validatePayload(t, data); !r.HasError(expecting) {
		t.Fatalf(testErrorExpected, expecting, r.Errors())
	}
}

func TestValidate_errorNotErrorObject(t *testing.T) {
	data := []byte(`{
	  "errors": [32]
	}`)

	if expecting, r := ErrNotErrorObject, validatePayload(t, data); !r.HasError(expecting) {
		t.Fatalf(testErrorExpected, expecting, r.Errors())
	}
}

func TestValidate_errorsKeys(t *testing.T) {
	data := []byte(`{
	  "errors": [{
			"aren55555": "foo"
		}]
	}`)

	if expecting, r := ErrInvalidErrorMember, validatePayload(t, data); !r.HasError(expecting) {
		t.Fatalf(testErrorExpected, expecting, r.Errors())
	}
}

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

	if validatePayload(t, data).HasErrors() {
		t.Fatal(testErrorNotExpected)
	}
}
