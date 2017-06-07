package test

import (
	"testing"

	"github.com/aren55555/jsonapivalidator"
)

func TestValidate_validAttributes(t *testing.T) {
	data := []byte(`{
	  "data": {
			"id": "1",
			"type": "car",
			"attributes": {
				"make":  "VW",
				"model": "R32",
				"year":  2008
			}
		}
	}`)

	expectedResult(t, data, noError, noWarning)
}

func TestValidate_validAttributes_nulls(t *testing.T) {
	data := []byte(`{
	  "data": {
			"id": "1",
			"type": "car",
			"attributes": {
				"make":  null,
				"model": null,
				"year":  null
			}
		}
	}`)

	expectedResult(t, data, noError, noWarning)
}

func TestValidate_invalidAttributes_topLevelRelationships(t *testing.T) {
	data := []byte(`{
	  "data": {
			"id": "1",
			"type": "car",
			"attributes": {
				"make":  "VW",
				"model": "R32",
				"year":  2008,
				"relationships": {}
			}
		}
	}`)

	expectedResult(t, data, noError, jsonapivalidator.WarnAttributesObjectHasRelationshipsMember)
}

func TestValidate_invalidAttributes_nestedRelationships(t *testing.T) {
	data := []byte(`{
	  "data": {
			"id": "1",
			"type": "car",
			"attributes": {
				"make":  "VW",
				"model": "R32",
				"year":  2008,
				"engine": [1, "abc", {
					"relationships": {}
				}]
			}
		}
	}`)

	expectedResult(t, data, noError, jsonapivalidator.WarnAttributesObjectHasRelationshipsMember)
}

func TestValidate_invalidAttributes_topLevelLinks(t *testing.T) {
	data := []byte(`{
	  "data": {
			"id": "1",
			"type": "car",
			"attributes": {
				"make":  "VW",
				"model": "R32",
				"year":  2008,
				"links": {}
			}
		}
	}`)

	expectedResult(t, data, noError, jsonapivalidator.WarnAttributesObjectHasLinksMember)
}

func TestValidate_invalidAttributes_nestedLinks(t *testing.T) {
	data := []byte(`{
	  "data": {
			"id": "1",
			"type": "car",
			"attributes": {
				"make":  "VW",
				"model": "R32",
				"year":  2008,
				"engine": [1, "abc", {
					"links": {}
				}]
			}
		}
	}`)

	expectedResult(t, data, noError, jsonapivalidator.WarnAttributesObjectHasLinksMember)
}
