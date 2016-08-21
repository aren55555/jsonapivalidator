package jsonapivalidator

func validateJSONAPIObject(ja interface{}, result *Result) {
	if _, ok := ja.(map[string]interface{}); !ok {
		result.AddError(ErrNotJSONAPIObject)
	}
}
