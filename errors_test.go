package jsonapivalidator

import "testing"

func TestValidate_errorsKeys(t *testing.T) {
	data := []byte(`{
	  "errors": {
			"aren55555": "foo"
		}
	}`)

	if !validatePayload(t, data).HasError(ErrInvalidErrorMember) {
		t.Fatal("Was expecting an error")
	}
}
