package test

import (
	"testing"

	"github.com/aren55555/jsonapivalidator"
)

func TestValidate_validResourceObject(t *testing.T) {
	data := []byte(`{
	  "data": {"id": "1", "type": "car"}
	}`)

	expectedResult(t, data, nil)
}

func TestValidate_validResourceIdentifierObject(t *testing.T) {
	data := []byte(`{
		"data": {"id": "1", "type": "car"}
	}`)

	expectedResult(t, data, nil)
}

func TestValidate_invalidResourceIdentifierObject_idNotString(t *testing.T) {
	data := []byte(`{
	  "data": {"id": [], "type": "car"}
	}`)

	expectedResult(t, data, jsonapivalidator.ErrIDNotString)
}

func TestValidate_validateResourceIdentifierObject_typeNotString(t *testing.T) {
	data := []byte(`{
	  "data": {"id": "1", "type": null}
	}`)

	expectedResult(t, data, jsonapivalidator.ErrTypeNotString)
}

func TestValidate_invalidResource(t *testing.T) {
	data := []byte(`{
	  "data": {"aren55555": true}
	}`)

	expectedResult(t, data, jsonapivalidator.ErrNotAResource)
}
