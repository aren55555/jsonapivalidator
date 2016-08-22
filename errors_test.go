package jsonapivalidator

import "testing"

func TestValidate_errorsNotArray(t *testing.T) {
	data := []byte(`{
	  "errors": 32
	}`)

	if !validatePayload(t, data).HasError(ErrInvalidErrorsType) {
		t.Fatal("Was expecting an error")
	}
}

func TestValidate_errorNotErrorObject(t *testing.T) {
	data := []byte(`{
	  "errors": [32]
	}`)

	if !validatePayload(t, data).HasError(ErrNotErrorObject) {
		t.Fatal("Was expecting an error")
	}
}

func TestValidate_errorsKeys(t *testing.T) {
	data := []byte(`{
	  "errors": [{
			"aren55555": "foo"
		}]
	}`)

	if !validatePayload(t, data).HasError(ErrInvalidErrorMember) {
		t.Fatal("Was expecting an error")
	}
}
