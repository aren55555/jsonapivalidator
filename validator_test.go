package jsonapivalidator

import (
	"encoding/json"
	"testing"
)

func TestValidate_ErrAtLeastOneRoot(t *testing.T) {
	data := []byte(`{}`)

	if !validatePayload(t, data).HasError(ErrAtLeastOneRoot) {
		t.Fatal("Was expecting an error")
	}
}

func TestValidate_ErrRootDataAndErrors(t *testing.T) {
	data := []byte(`{
  "data": {},
  "errors": {}
  }`)

	if !validatePayload(t, data).HasError(ErrRootDataAndErrors) {
		t.Fatal("Was expecting an error")
	}
}

func TestValidate_nullData(t *testing.T) {
	data := []byte(`{
  "data": null
  }`)

	if validatePayload(t, data).HasErrors() {
		t.Fatal("Was not expecting an error")
	}
}

func TestValidate_errorsKeys(t *testing.T) {
	data := []byte(`{
	  "errors": {
			"aren55555": "foo"
		}
	  }`)

	if !validatePayload(t, data).HasError(ErrInvalidErrorMember) {
		t.Fatal("Was expecting an error")
	}
}

func validatePayload(t *testing.T, data []byte) *Result {
	var obj map[string]interface{}
	if err := json.Unmarshal(data, &obj); err != nil {
		t.Fatal(err)
	}
	return Validate(obj)
}
