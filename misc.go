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

func validateURL(t interface{}, result *Result) {
	// TODO: figure out how to validate a URL; url.Parse will accept almost any
	// string without producing an error
}
