package jsonapivalidator

import "encoding/json"

// UnmarshalAndValidate wil
func UnmarshalAndValidate(data []byte) (result *Result, err error) {
	var root interface{}

	if err = json.Unmarshal(data, &root); err != nil {
		return
	}

	result = Validate(root)

	return
}

// Validate checks the root payload against the JSONAPI sepc, returns a Result
// populated with all spec violations.
func Validate(root interface{}) (result *Result) {
	result = NewResult()

	// Check the document
	document := root.(map[string]interface{})
	data, dataExists := document[memberData]
	errors, errorsExists := document[memberErrors]
	meta, metaExists := document[memberMeta]

	if !(metaExists || errorsExists || dataExists) {
		result.AddError(ErrAtLeastOneRoot)
	}

	if dataExists && errorsExists {
		result.AddError(ErrRootDataAndErrors)
	}

	// Validate /data
	if dataExists {
		validateData(data, result)
	}

	// Validate /errors
	if errorsExists {
		validateErrors(errors, result)
	}

	// Validate /meta
	if metaExists {
		validateMetaObject(meta, result)
	}

	// A document MAY contain any of these top-level members:
	//  jsonapi: an object describing the server’s implementation
	if jsonapi, jsonAPIExists := document[memberJSONAPI]; jsonAPIExists {
		validateJSONAPIObject(jsonapi, result)
	}
	//  links: a links object related to the primary data.
	if links, linksExists := document[memberLinks]; linksExists {
		validateLinksObject(links, result, linksAllMembers)
	}
	//  included: an array of resource objects that are related to the primary
	//            data and/or each other (“included resources”).
	included, includedExists := document[memberIncluded]
	// If a document does not contain a top-level data key, the included member
	// MUST NOT be present either.
	if !dataExists && includedExists {
		result.AddError(ErrRootIncludedWithoutData)
	}
	if includedExists {
		validateIncluded(included, result)
	}

	return
}
