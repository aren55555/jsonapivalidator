package test

import (
	"testing"

	"github.com/aren55555/jsonapivalidator"
)

func TestValidate_nullData(t *testing.T) {
	data := []byte(`{
  	"data": null
  }`)

	expectedResult(t, data, nil)
}

func TestValidate_dataArrayEmpty(t *testing.T) {
	data := []byte(`{
  	"data": []
  }`)

	expectedResult(t, data, nil)
}

func TestValidate_dataUnexpected(t *testing.T) {
	data := []byte(`{
	  "data": false
	}`)

	expectedResult(t, data, jsonapivalidator.ErrInvalidDataType)
}
