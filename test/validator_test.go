package test

import (
	"testing"

	"github.com/aren55555/jsonapivalidator"
)

func TestValidate_ErrAtLeastOneRoot(t *testing.T) {
	data := []byte(`{}`)

	if expecting, r := jsonapivalidator.ErrAtLeastOneRoot, validatePayload(t, data); !r.HasError(expecting) {
		t.Fatalf(testErrorExpected, expecting, r.Errors())
	}
}

func TestValidate_ErrRootDataAndErrors(t *testing.T) {
	data := []byte(`{
  	"data": {},
  	"errors": {}
  }`)

	if expecting, r := jsonapivalidator.ErrRootDataAndErrors, validatePayload(t, data); !r.HasError(expecting) {
		t.Fatalf(testErrorExpected, expecting, r.Errors())
	}
}

// func TestValidate_ErrRootDataAndErrors(t *testing.T) {
// 	data := []byte(`{
//   	"data": {},
//   	"errors": {}
//   }`)
//
// 	if expecting, r := jsonapivalidator.ErrRootDataAndErrors, validatePayload(t, data); !r.HasError(expecting) {
// 		t.Fatalf(testErrorExpected, expecting, r.Errors())
// 	}
// }
