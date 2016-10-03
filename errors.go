package jsonapivalidator

func validateErrors(e interface{}, result *Result) {
	errors, ok := e.([]interface{})
	if !ok {
		result.AddError(ErrInvalidErrorsType)
		return // cannot procceed
	}

	// Validate each element is an error object
	for _, err := range errors {
		validateErrorObject(err, result)
	}
}

// errors map[string]interface{}, result *Result
func validateErrorObject(e interface{}, result *Result) {
	// The value of the attributes key MUST be an object (an “attributes object”)
	err, ok := e.(map[string]interface{})
	if !ok {
		result.AddError(ErrNotErrorObject)
		return // cannot procceed
	}

	for k, v := range err {
		switch k {
		case memberID:
			// a unique identifier for this particular occurrence of the problem
			// TODO: implement
		case memberLinks:
			// a links object containing the following members:
			//  about: a link that leads to further details about this particular
			//         occurrence of the problem.
			validateLinksObject(v, result, nil)
		case memberStatus:
			// the HTTP status code applicable to this problem, expressed as a string
			// value.
			// TODO: implement
		case memberCode:
			// an application-specific error code, expressed as a string value
			// TODO: implement
		case memberTitle:
			// a short, human-readable summary of the problem that SHOULD NOT change
			// from occurrence to occurrence of the problem, except for purposes of
			// localization.
			// TODO: implement
		case memberDetail:
			// a human-readable explanation specific to this occurrence of the
			// problem. Like title, this field’s value can be localized
			// TODO: implement
		case memberSource:
			// an object containing references to the source of the error, optionally
			// including any of the following members:
			//  pointer: a JSON Pointer [RFC6901] to the associated entity in the
			//           request document [e.g. "/data" for a primary data object, or
			//           "/data/attributes/title" for a specific attribute]
			//  parameter: a string indicating which URI query parameter caused the
			//             error
			// TODO: implement
		case memberMeta:
			// a meta object containing non-standard meta-information about the error
			validateMetaObject(v, result)
		default:
			result.AddError(ErrInvalidErrorMember)
		}
	}
}
