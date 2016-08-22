package test

import (
	"testing"

	"github.com/aren55555/jsonapivalidator"
)

func TestValidate_ErrAtLeastOneRoot(t *testing.T) {
	data := []byte(`{}`)
	expectedResult(t, data, jsonapivalidator.ErrAtLeastOneRoot)
}

func TestValidate_ErrRootDataAndErrors(t *testing.T) {
	data := []byte(`{
  	"data": {},
  	"errors": {}
  }`)
	expectedResult(t, data, jsonapivalidator.ErrRootDataAndErrors)
}

func TestValidate_invalidIncludedWithoutData(t *testing.T) {
	data := []byte(`{
  	"included": []
  }`)
	expectedResult(t, data, jsonapivalidator.ErrRootIncludedWithoutData)
}
