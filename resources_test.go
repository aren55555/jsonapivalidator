package jsonapivalidator

import "testing"

func TestValidate_validateResourceIdentifierObject_idNotString(t *testing.T) {
	data := []byte(`{
	  "data": {"id": [], "type": "car"}
	}`)

	if expecting, r := ErrIDNotString, validatePayload(t, data); !r.HasError(expecting) {
		t.Fatalf(testErrorExpected, expecting, r.Errors())
	}
}

func TestValidate_validateResourceIdentifierObject_typeNotString(t *testing.T) {
	data := []byte(`{
	  "data": {"id": "1", "type": null}
	}`)

	if expecting, r := ErrTypeNotString, validatePayload(t, data); !r.HasError(expecting) {
		t.Fatalf(testErrorExpected, expecting, r.Errors())
	}
}

func TestValidate_validateResourceObject(t *testing.T) {
	data := []byte(`{
	  "data": {"id": "1", "type": "car"}
	}`)

	if validatePayload(t, data).HasErrors() {
		t.Fatal(testErrorNotExpected)
	}
}

func TestValidate_invalidResource(t *testing.T) {
	data := []byte(`{
	  "data": {"aren55555": true}
	}`)

	if expecting, r := ErrNotAResource, validatePayload(t, data); !r.HasError(expecting) {
		t.Fatalf(testErrorExpected, expecting, r.Errors())
	}
}

func TestValidate_validateResourceIdentifierObject(t *testing.T) {
	data := []byte(`{
		"data": {"id": "1", "type": "car"}
	}`)

	if validatePayload(t, data).HasErrors() {
		t.Fatal(testErrorNotExpected)
	}
}
