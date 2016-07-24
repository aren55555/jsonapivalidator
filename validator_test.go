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

	if expecting, r := ErrRootDataAndErrors, validatePayload(t, data); !r.HasError(expecting) {
		t.Fatalf("Was expecting an error\nExpected: %s\nGot: %s", expecting, r.Errors())
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

func TestValidate_dataArrayEmpty(t *testing.T) {
	data := []byte(`{
  "data": []
  }`)

	if validatePayload(t, data).HasErrors() {
		t.Fatal("Was not expecting an error")
	}
}

func TestValidate_dataUnexpected(t *testing.T) {
	data := []byte(`{
	  "data": false
	  }`)

	if expecting, r := ErrInvalidDataType, validatePayload(t, data); !r.HasError(expecting) {
		t.Fatalf("Was expecting an error\nExpected: %s\nGot: %s", expecting, r.Errors())
	}
}

func TestValidate_validateResourceObject(t *testing.T) {
	data := []byte(`{
	  "data": {"id": "1", "type": "car"}
	  }`)

	if validatePayload(t, data).HasErrors() {
		t.Fatal("Was not expecting an error")
	}
}

func TestValidate_invalidResource(t *testing.T) {
	data := []byte(`{
	  "data": {"aren55555": true}
	  }`)

	if expecting, r := ErrNotAResource, validatePayload(t, data); !r.HasError(expecting) {
		t.Fatalf("Was expecting an error\nExpected: %s\nGot: %s", expecting, r.Errors())
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
