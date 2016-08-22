package test

import (
	"encoding/json"
	"testing"

	"github.com/aren55555/jsonapivalidator"
)

const (
	testErrorExpected    = "Was expecting an error\nExpected: %s\nGot: %s"
	testErrorNotExpected = "Was not expecting an error"
)

func validatePayload(t *testing.T, data []byte) *jsonapivalidator.Result {
	var obj map[string]interface{}
	if err := json.Unmarshal(data, &obj); err != nil {
		t.Fatal(err)
	}
	return jsonapivalidator.Validate(obj)
}

func expectedResult(t *testing.T, data []byte, expectedErr error) {
	r := validatePayload(t, data)

	if expectedErr == nil {
		// Was not expecting errors
		if r.HasErrors() {
			t.Fatal(testErrorNotExpected)
		}
		return
	}

	// An error is expected
	if !r.HasError(expectedErr) {
		t.Fatalf(testErrorExpected, expectedErr, r.Errors())
	}
}
