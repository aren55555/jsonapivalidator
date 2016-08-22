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

func TestUnmarshalAndValidate(t *testing.T) {
	data := loadSample(t, "default.json")

	r, err := jsonapivalidator.UnmarshalAndValidate(data)
	if err != nil {
		t.Fatal(err)
	}

	if len(r.Errors()) > 0 {
		t.Fatal("Was not expecting any errors")
	}
}
