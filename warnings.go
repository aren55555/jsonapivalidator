package jsonapivalidator

import "errors"

var (
	// WarnAttributesObjectFK is the warning produced when FKs may have been
	// detected in the attributes object
	WarnAttributesObjectFK = errors.New("Although has-one foreign keys (e.g. author_id) " +
		"are often stored internally alongside other information to be " +
		"represented in a resource object, these keys SHOULD NOT appear as " +
		"attributes.")

	// WarnAttributesObjectHasRelationshipsMember is the warning produced when a
	// relationships member name has been detected anywhere in the attributes
	// object
	WarnAttributesObjectHasRelationshipsMember = errors.New("Complex data " +
		"structures involving JSON objects and arrays are allowed as attribute " +
		"values. However, any object that constitutes or is contained in an " +
		"attribute MUST NOT contain a relationships member, as it has been " +
		"reserved by this specification for future use.")

	// WarnAttributesObjectHasLinksMember is the warning produced when a
	// relationships member name has been detected anywhere in the attributes
	// object
	WarnAttributesObjectHasLinksMember = errors.New("Complex data " +
		"structures involving JSON objects and arrays are allowed as attribute " +
		"values. However, any object that constitutes or is contained in an " +
		"attribute MUST NOT contain a links member, as it has been " +
		"reserved by this specification for future use.")
)
