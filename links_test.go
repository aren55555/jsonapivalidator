package jsonapivalidator

import "testing"

func TestValidate_validLinks(t *testing.T) {
	data := []byte(`{
		"meta": {},
		"links": {
			"self": "http://example.com/articles?page[number]=3&page[size]=1",
    	"first": "http://example.com/articles?page[number]=1&page[size]=1",
    	"prev": "http://example.com/articles?page[number]=2&page[size]=1",
    	"next": "http://example.com/articles?page[number]=4&page[size]=1",
    	"last": "http://example.com/articles?page[number]=13&page[size]=1",
  		"related": {
    		"href": "http://example.com/articles/1/comments",
    		"meta": {
      		"count": 10
    		}
  		}
		}
	}`)

	if validatePayload(t, data).HasErrors() {
		t.Fatal(testErrorNotExpected)
	}
}

func TestValidate_invalidLinks(t *testing.T) {
	data := []byte(`{
		"meta": {},
	  "links": 5
	}`)

	if expecting, r := ErrNotLinksObject, validatePayload(t, data); !r.HasError(expecting) {
		t.Fatalf(testErrorExpected, expecting, r.Errors())
	}
}

func TestValidate_invalidLinkValue(t *testing.T) {
	data := []byte(`{
		"meta": {},
	  "links": {"aren": []}
	}`)

	if expecting, r := ErrInvalidLinkType, validatePayload(t, data); !r.HasError(expecting) {
		t.Fatalf(testErrorExpected, expecting, r.Errors())
	}
}

func TestValidate_invalidLinkObjectMember(t *testing.T) {
	data := []byte(`{
		"meta": {},
	  "links": {
			"aren": {"foo": "bar"}
		}
	}`)

	if expecting, r := ErrInvalidLinkMember, validatePayload(t, data); !r.HasError(expecting) {
		t.Fatalf(testErrorExpected, expecting, r.Errors())
	}
}
