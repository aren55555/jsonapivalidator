package test

import (
	"testing"

	"github.com/aren55555/jsonapivalidator"
)

func TestValidate_validResourceObject(t *testing.T) {
	data := []byte(`{
	  "data": {"id": "1", "type": "car"}
	}`)

	expectedResult(t, data, noError, noWarning)
}

func TestValidate_validResourceIdentifierObject(t *testing.T) {
	data := []byte(`{
		"data": {"id": "1", "type": "car"}
	}`)

	expectedResult(t, data, noError, noWarning)
}

func TestValidate_invalidResourceIdentifierObject_idNotString(t *testing.T) {
	data := []byte(`{
	  "data": {"id": [], "type": "car"}
	}`)

	expectedResult(t, data, jsonapivalidator.ErrIDNotString, noWarning)
}

func TestValidate_validateResourceIdentifierObject_typeNotString(t *testing.T) {
	data := []byte(`{
	  "data": {"id": "1", "type": null}
	}`)

	expectedResult(t, data, jsonapivalidator.ErrTypeNotString, noWarning)
}

func TestValidate_invalidResource(t *testing.T) {
	data := []byte(`{
	  "data": {"aren55555": true}
	}`)

	expectedResult(t, data, jsonapivalidator.ErrNotAResource, noWarning)
}
