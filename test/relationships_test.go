package test

import (
	"testing"

	"github.com/aren55555/jsonapivalidator"
)

func TestValidate_invalidRelationshipsObject(t *testing.T) {
	data := []byte(`{
	  "data": {
			"id": "1",
			"type": "car",
			"relationships": [1,2,3]
		}
	}`)

	if expecting, r := jsonapivalidator.ErrNotRelationshipsObject, validatePayload(t, data); !r.HasError(expecting) {
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

	if expecting, r := jsonapivalidator.ErrNotRelationshipObject, validatePayload(t, data); !r.HasError(expecting) {
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

	if expecting, r := jsonapivalidator.ErrInvalidResourceLinkage, validatePayload(t, data); !r.HasError(expecting) {
		t.Fatalf(testErrorExpected, expecting, r.Errors())
	}
}
