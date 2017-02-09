package test

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/aren55555/jsonapivalidator"
)

var (
	noError, noWarning error
)

const (
	testErrorExpected       = "Was expecting an error\nExpected: %s\nGot: %s"
	testErrorLengthExpected = "Was expecting %d errors; got %d\nErrors: %s"
	testErrorNotExpected    = "Was not expecting an error\nGot: %s"

	testWarningExpected       = "Was expecting a warning\nExpected: %s\nGot: %s"
	testWarningLengthExpected = "Was expecting %d warnings; got %d\nWarnings: %s"
	testWarningNotExpected    = "Was not expecting a warning\nGot: %s"
)

func validatePayload(t *testing.T, data []byte) *jsonapivalidator.Result {
	var obj interface{}
	if err := json.Unmarshal(data, &obj); err != nil {
		t.Fatal(err)
	}
	return jsonapivalidator.Validate(obj)
}

func expectedResult(t *testing.T, data []byte, expectedErr error, expectedWarning error) {
	r := validatePayload(t, data)

	if expectedErr == nil {
		// Was not expecting errors
		if r.HasErrors() {
			t.Fatalf(testErrorNotExpected, joinErrors(r.Errors()))
		}
	} else {
		// An error is expected
		if !r.HasError(expectedErr) {
			t.Fatalf(testErrorExpected, expectedErr, joinErrors(r.Errors()))
		}
	}

	if expectedWarning == nil {
		// Was not expecting warnings
		if r.HasWarnings() {
			t.Fatalf(testWarningNotExpected, joinErrors(r.Warnings()))
		}
	} else {
		// A warning is expected
		if !r.HasWarning(expectedWarning) {
			t.Fatalf(testWarningExpected, expectedWarning, joinErrors(r.Warnings()))
		}
	}
}

func expectedResultHasErrors(t *testing.T, data []byte, errorCount uint) {
	r := validatePayload(t, data)

	if a := uint(len(r.Errors())); a != errorCount {
		t.Fatalf(testErrorLengthExpected, errorCount, a, joinErrors(r.Errors()))
	}
}

func loadSample(t *testing.T, sample string) (data io.Reader) {
	sampleFile, err := filepath.Abs(fmt.Sprintf("./samples/%s", sample))
	if err != nil {
		t.Fatal(err)
	}

	data, err = os.Open(sampleFile)
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
