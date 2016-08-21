package jsonapivalidator

func validateAttributesObject(a interface{}, result *Result) {
	// The value of the attributes key MUST be an object (an “attributes object”).
	_, ok := a.(map[string]interface{})
	if !ok {
		result.AddError(ErrNotAttributesObject)
		return // cannot procceed
	}

	// TODO:
	// Complex data structures involving JSON objects and arrays are allowed as
	// attribute values. However, any object that constitutes or is contained in
	// an attribute MUST NOT contain a relationships or links member, as those
	// members are reserved by this specification for future use.

	// TODO:
	// Although has-one foreign keys (e.g. author_id) are often stored internally
	// alongside other information to be represented in a resource object, these
	// keys SHOULD NOT appear as attributes.
}
