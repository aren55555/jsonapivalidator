package jsonapivalidator

func validateIncluded(included interface{}, result *Result) {
	// included MUST be an array
	includes, ok := included.([]interface{})
	if !ok {
		result.AddError(ErrInvalidIncludedType)
	}

	for _, i := range includes {
		r := i.(map[string]interface{})
		validateResourceObject(r, result)
	}
}
