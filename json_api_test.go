package jsonapivalidator

import "testing"

func TestValidate_invalidJSONAPI(t *testing.T) {
	data := []byte(`{
  	"data": {},
  	"jsonapi": [1,2,3]
  }`)

	if expecting, r := ErrNotJSONAPIObject, validatePayload(t, data); !r.HasError(expecting) {
		t.Fatalf(testErrorExpected, expecting, r.Errors())
	}
}
