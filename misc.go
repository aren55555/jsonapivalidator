package jsonapivalidator

func validateID(id interface{}, result *Result) {
	if _, ok := id.(string); !ok {
		result.AddError(ErrIDNotString)
	}
	return
}

func validateType(t interface{}, result *Result) {
	if _, ok := t.(string); !ok {
		result.AddError(ErrTypeNotString)
	}
	return
}

func validateJSONAPIObject(ja interface{}, result *Result) {
	if _, ok := ja.(map[string]interface{}); !ok {
		result.AddError(ErrNotJSONAPIObject)
	}
}
