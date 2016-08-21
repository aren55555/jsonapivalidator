package jsonapivalidator

import "testing"

func TestValidate_nullData(t *testing.T) {
	data := []byte(`{
  	"data": null
  }`)

	if validatePayload(t, data).HasErrors() {
		t.Fatal(testErrorNotExpected)
	}
}

func TestValidate_dataArrayEmpty(t *testing.T) {
	data := []byte(`{
  	"data": []
  }`)

	if validatePayload(t, data).HasErrors() {
		t.Fatal(testErrorNotExpected)
	}
}

func TestValidate_dataUnexpected(t *testing.T) {
	data := []byte(`{
	  "data": false
	}`)

	if expecting, r := ErrInvalidDataType, validatePayload(t, data); !r.HasError(expecting) {
		t.Fatalf(testErrorExpected, expecting, r.Errors())
	}
}
