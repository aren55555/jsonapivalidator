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

func TestValidate_invalidJSONAPI(t *testing.T) {
	data := []byte(`{
  	"data": {},
  	"jsonapi": [1,2,3]
  }`)

	if expecting, r := ErrNotJSONAPIObject, validatePayload(t, data); !r.HasError(expecting) {
		t.Fatalf(testErrorExpected, expecting, r.Errors())
	}
}

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

func TestValidate_validateesourceIdentifierObject(t *testing.T) {
	data := []byte(`{
		"data": {"id": "1", "type": "car"}
	}`)

	if validatePayload(t, data).HasErrors() {
		t.Fatal(testErrorNotExpected)
	}
}

func TestValidate_validateesourceIdentifierObject_idNotString(t *testing.T) {
	data := []byte(`{
	  "data": {"id": [], "type": "car"}
	}`)

	if expecting, r := ErrIDNotString, validatePayload(t, data); !r.HasError(expecting) {
		t.Fatalf(testErrorExpected, expecting, r.Errors())
	}
}

func TestValidate_validateesourceIdentifierObject_typeNotString(t *testing.T) {
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

func TestValidate_attributes(t *testing.T) {
	data := []byte(`{
	  "data": {
			"id": "1",
			"type": "car",
			"attributes": {
				"make":  "VW",
				"model": "R32",
				"year":  2008
			}
		}
	}`)

	if validatePayload(t, data).HasErrors() {
		t.Fatal(testErrorNotExpected)
	}
}

func TestValidate_invalidRelationshipsObject(t *testing.T) {
	data := []byte(`{
	  "data": {
			"id": "1",
			"type": "car",
			"relationships": [1,2,3]
		}
	}`)

	if expecting, r := ErrNotRelationshipsObject, validatePayload(t, data); !r.HasError(expecting) {
		t.Fatalf(testErrorExpected, expecting, r.Errors())
	}
}

func TestValidate_invalidRelationshipObject(t *testing.T) {
	data := []byte(`{
	  "data": {
			"id": "1",
			"type": "car",
			"relationships": {
				"driver": 55555
			}
		}
	}`)

	if expecting, r := ErrNotRelationshipObject, validatePayload(t, data); !r.HasError(expecting) {
		t.Fatalf(testErrorExpected, expecting, r.Errors())
	}
}

func TestValidate_nullResourceLinkage(t *testing.T) {
	data := []byte(`{
	  "data": {
			"id": "1",
			"type": "car",
			"relationships": {
				"driver": {
					"data": null
				}
			}
		}
	}`)

	if validatePayload(t, data).HasErrors() {
		t.Fatal(testErrorNotExpected)
	}
}

func TestValidate_invalidResourceLinkage(t *testing.T) {
	data := []byte(`{
	  "data": {
			"id": "1",
			"type": "car",
			"relationships": {
				"driver": {
					"data": 42
				}
			}
		}
	}`)

	if expecting, r := ErrInvalidResourceLinkage, validatePayload(t, data); !r.HasError(expecting) {
		t.Fatalf(testErrorExpected, expecting, r.Errors())
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
