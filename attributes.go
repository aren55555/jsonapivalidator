package jsonapivalidator

import (
	"reflect"
	"strings"
)

func validateAttributesObject(a interface{}, result *Result) {
	// The value of the attributes key MUST be an object (an “attributes object”).
	attributes, ok := a.(map[string]interface{})
	if !ok {
		result.AddError(ErrNotAttributesObject)
		return // cannot procceed
	}

	for k, v := range attributes {
		validateAttributeMemberName(k, result)
		validateAttributeValue(v, result)
	}
}

func validateAttributeMemberName(m string, result *Result) {
	// Although has-one foreign keys (e.g. author_id) are often stored internally
	// alongside other information to be represented in a resource object, these
	// keys SHOULD NOT appear as attributes.
	if strings.Contains(strings.ToLower(m), "id") {
		// TODO: improve this with a regex; currently it will product false warnings
		//       for words that contain "id"
		result.AddWarning(WarnAttributesObjectFK)
	}

	// Complex data structures involving JSON objects and arrays are allowed as
	// attribute values. However, any object that constitutes or is contained in
	// an attribute MUST NOT contain a relationships or links member, as those
	// members are reserved by this specification for future use.
	if m == memberRelationships {
		result.AddWarning(WarnAttributesObjectHasRelationshipsMember)
	}
	if m == memberLinks {
		result.AddWarning(WarnAttributesObjectHasLinksMember)
	}
}

func validateAttributeValue(v interface{}, result *Result) {
	// Complex data structures involving JSON objects and arrays are allowed as
	// attribute values. However, any object that constitutes or is contained in
	// an attribute MUST NOT contain a relationships or links member, as those
	// members are reserved by this specification for future use.

	if v == nil {
		return // JSON attribute was null
	}

	// We had a non null value
	switch reflect.TypeOf(v).Kind() {
	case reflect.Map:
		// TODO: check this map for relationships or links members
		mapValue, ok := v.(map[string]interface{})
		if !ok {
			result.AddError(ErrNotLinkObject) // TODO: what error?
			return
		}

		for k, v := range mapValue {
			validateAttributeMemberName(k, result)
			validateAttributeValue(v, result)
		}
	case reflect.Slice:
		// Check each element of this slice for relationships or links members
		sliceValues, ok := v.([]interface{})
		if !ok {
			result.AddError(ErrNotLinkObject) // TODO: what error?
			return
		}

		for _, sv := range sliceValues {
			validateAttributeValue(sv, result)
		}
	default:
		return
	}
}
