package test

import "testing"

func TestValidate_attributes(t *testing.T) {
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

	expectedResult(t, data, nil)
}
