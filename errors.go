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
	// The value of the attributes key MUST be an object (an “attributes object”).
	err, ok := e.(map[string]interface{})
	if !ok {
		result.AddError(ErrNotErrorObject)
		return // cannot procceed
	}

	for k := range err {
		switch k {
		case memberID:
		case memberLinks:
		case memberStatus:
		case memberCode:
		case memberTitle:
		case memberDetail:
		case memberSource:
		case memberMeta:
		default:
			result.AddError(ErrInvalidErrorMember)
		}
	}
}
