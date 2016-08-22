package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/aren55555/jsonapivalidator"
)

const (
	testErrorExpected       = "Was expecting an error\nExpected: %s\nGot: %s"
	testErrorLengthExpected = "Was expecting %d errors; got %d\nErrors: %s"
	testErrorNotExpected    = "Was not expecting an error\nGot: %s"
)

func validatePayload(t *testing.T, data []byte) *jsonapivalidator.Result {
	var obj interface{}
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
			t.Fatalf(testErrorNotExpected, joinErrors(r.Errors()))
		}
		return
	}

	// An error is expected
	if !r.HasError(expectedErr) {
		t.Fatalf(testErrorExpected, expectedErr, joinErrors(r.Errors()))
	}
}

func expectedResultHasErrors(t *testing.T, data []byte, errorCount uint) {
	r := validatePayload(t, data)

	if a := uint(len(r.Errors())); a != errorCount {
		t.Fatalf(testErrorLengthExpected, errorCount, a, joinErrors(r.Errors()))
	}
}

func loadSample(t *testing.T, sample string) (data []byte) {
	sampleFile, err := filepath.Abs(fmt.Sprintf("./samples/%s", sample))
	if err != nil {
		t.Fatal(err)
	}

	data, err = ioutil.ReadFile(sampleFile)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func joinErrors(errors []error) (r string) {
	strs := []string{}
	for _, err := range errors {
		strs = append(strs, fmt.Sprintf(">>> %s <<<", err))
	}

	return strings.Join(strs, ",\n")
}
