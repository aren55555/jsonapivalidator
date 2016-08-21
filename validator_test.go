package jsonapivalidator

import (
	"encoding/json"
	"testing"
)

const (
	testErrorExpected    = "Was expecting an error\nExpected: %s\nGot: %s"
	testErrorNotExpected = "Was not expecting an error"
)

func TestValidate_ErrAtLeastOneRoot(t *testing.T) {
	data := []byte(`{}`)

	if expecting, r := ErrAtLeastOneRoot, validatePayload(t, data); !r.HasError(expecting) {
		t.Fatalf(testErrorExpected, expecting, r.Errors())
	}
}

func TestValidate_ErrRootDataAndErrors(t *testing.T) {
	data := []byte(`{
  	"data": {},
  	"errors": {}
  }`)

	if expecting, r := ErrRootDataAndErrors, validatePayload(t, data); !r.HasError(expecting) {
		t.Fatalf(testErrorExpected, expecting, r.Errors())
	}
}

func validatePayload(t *testing.T, data []byte) *Result {
	var obj map[string]interface{}
	if err := json.Unmarshal(data, &obj); err != nil {
		t.Fatal(err)
	}
	return Validate(obj)
}
