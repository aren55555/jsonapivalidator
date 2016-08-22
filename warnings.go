package jsonapivalidator

const (
	// WarnAttributesObjectFK is the warning produced when FKs may have been
	// detected in the attributes object
	WarnAttributesObjectFK = "Although has-one foreign keys (e.g. author_id) " +
		"are often stored internally alongside other information to be " +
		"represented in a resource object, these keys SHOULD NOT appear as " +
		"attributes."
)
