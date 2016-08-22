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
