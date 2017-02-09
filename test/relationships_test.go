package test

import (
	"testing"

	"github.com/aren55555/jsonapivalidator"
)

func TestValidate_invalidRelationshipsObject(t *testing.T) {
	data := []byte(`{
	  "data": {
			"id": "1",
			"type": "car",
			"relationships": [1,2,3]
		}
	}`)

	expectedResult(t, data, jsonapivalidator.ErrNotRelationshipsObject)
}

func TestValidate_invalidRelationshipObject(t *testing.T) {
	data := []byte(`{
	  "data": {
			"id": "1",
			"type": "car",
			"relationships": {
				"driver": 55555
			}
		}
	}`)

	expectedResult(t, data, jsonapivalidator.ErrNotRelationshipObject)
}

func TestValidate_nullResourceLinkage(t *testing.T) {
	data := []byte(`{
	  "data": {
			"id": "1",
			"type": "car",
			"relationships": {
				"driver": {
					"data": null
				}
			}
		}
	}`)

	expectedResult(t, data, nil)
}

func TestValidate_invalidResourceLinkage(t *testing.T) {
	data := []byte(`{
	  "data": {
			"id": "1",
			"type": "car",
			"relationships": {
				"driver": {
					"data": 42
				}
			}
		}
	}`)

	expectedResult(t, data, jsonapivalidator.ErrInvalidResourceLinkage)
}

func TestValidate_validToManyRelation(t *testing.T) {
	data := []byte(`{
	  "data": {
			"id": "1",
			"type": "car",
			"relationships": {
				"drivers": {
					"data": [{
						"type": "person",
						"id": "1"
					}, {
						"type": "person",
						"id": "2"
					}]
				}
			}
		}
	}`)

	expectedResult(t, data, nil)
}

func TestValidate_validToManyRelation_withValidLinks(t *testing.T) {
	data := []byte(`{
	  "data": {
			"id": "1",
			"type": "car",
			"relationships": {
				"drivers": {
					"links": {
						"self": "http://example.com/cars/1/relationships/drivers"
					},
					"data": [{
						"type": "person",
						"id": "1"
					}, {
						"type": "person",
						"id": "2"
					}]
				}
			}
		}
	}`)

	expectedResult(t, data, nil)
}

func TestValidate_validToManyRelation_withInvalidLinks(t *testing.T) {
	data := []byte(`{
	  "data": {
			"id": "1",
			"type": "car",
			"relationships": {
				"drivers": {
					"links": {
						"self": []
					},
					"data": [{
						"type": "person",
						"id": "1"
					}, {
						"type": "person",
						"id": "2"
					}]
				}
			}
		}
	}`)

	expectedResult(t, data, jsonapivalidator.ErrInvalidLinkType)
}

func TestValidate_validToOneRelation(t *testing.T) {
	data := []byte(`{
	  "data": {
			"id": "1",
			"type": "car",
			"relationships": {
				"driver": {
					"data": {
						"type": "person",
						"id": "1"
					}
				}
			}
		}
	}`)

	expectedResult(t, data, nil)
}

func TestValidate_validToOneRelation_withValidLinks(t *testing.T) {
	data := []byte(`{
	  "data": {
			"id": "1",
			"type": "car",
			"relationships": {
				"driver": {
					"links": {
						"self": "http://example.com/cars/1/relationships/driver"
					},
					"data": {
						"type": "person",
						"id": "1"
					}
				}
			}
		}
	}`)

	expectedResult(t, data, nil)
}

func TestValidate_validToOneRelation_withInvalidLinks(t *testing.T) {
	data := []byte(`{
	  "data": {
			"id": "1",
			"type": "car",
			"relationships": {
				"driver": {
					"links": {
						"self": 5
					},
					"data": {
						"type": "person",
						"id": "1"
					}
				}
			}
		}
	}`)

	expectedResult(t, data, jsonapivalidator.ErrInvalidLinkType)
}
