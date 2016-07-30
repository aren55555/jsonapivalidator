package jsonapivalidator

import (
	"fmt"
	"reflect"
)

// Validate checks the payload against the JSONAPI sepc, returns a Result
// populated with all spec violations.
func Validate(payload interface{}) (result *Result) {
	result = NewResult()

	// Check the document
	document := payload.(map[string]interface{})
	data, dataExists := document[memberData]
	errors, errorsExists := document[memberErrors]
	_, metaExists := document[memberMeta]

	if !(metaExists || errorsExists || dataExists) {
		result.AddError(ErrAtLeastOneRoot)
	}

	if dataExists && errorsExists {
		result.AddError(ErrRootDataAndErrors)
	}

	// Validate /data
	if dataExists {
		validateData(data, result)
	}

	// Validate /errors
	if errorsExists {
		e := errors.(map[string]interface{})
		validateErrors(e, result)
	}

	// A document MAY contain any of these top-level members: "jsonapi", "links",
	// "included"
	if jsonapi, jsonAPIExists := document[memberJSONAPI]; jsonAPIExists {
		validateJSONAPIObject(jsonapi, result)
	}

	return
}

func validateData(data interface{}, result *Result) {
	if data == nil {
		return // "data": null is an allowed value
	}

	switch reflect.TypeOf(data).Kind() {
	case reflect.Map:
		// Single
		r := data.(map[string]interface{})
		validateResource(r, result)
	case reflect.Slice:
		// Array
		resources := data.([]interface{})
		for _, resource := range resources {
			r := resource.(map[string]interface{})
			validateResource(r, result)
		}
	default:
		result.AddError(ErrInvalidDataType)
	}
}

func validateResource(r map[string]interface{}, result *Result) {
	// We may have either a RIO or RO
	// RO = Resource Objec
	// RIO = Resource Identifier Object

	isRO := isResourceObject(r)
	isRIO := isResourceIdentifierObject(r)

	if !isRO && !isRIO {
		result.AddError(ErrNotAResource)
		return
	}

	if isRO {
		validateResourceObject(r, result)
		return // return since RO is a superset of RIO
	}

	validateResourceIdentifierObject(r, result)
	return
}

func isResourceObject(r map[string]interface{}) bool {
	for key := range r {
		switch key {
		case memberID, memberType, memberAttributes, memberRelationships, memberLinks, memberMeta:
			// do nothing
		default:
			return false
		}
	}
	return true
}

func validateResourceObject(ro map[string]interface{}, result *Result) {
	// A resource object MUST contain at least the following top-level members:
	if id, hasID := ro[memberID]; !hasID {
		// TODO: The id member is not required when the resource object originates
		// at the client and represents a new resource to be created on the server.
		result.AddError(ErrResourceObjectMissingID)
	} else {
		validateID(id, result)
	}
	if jType, hasType := ro[memberType]; !hasType {
		result.AddError(ErrResourceObjectMissingType)
	} else {
		validateType(jType, result)
	}

	// A resource object MAY contain any of these top-level members:
	if attributes, hasAttributes := ro[memberAttributes]; hasAttributes {
		validateAttributesObject(attributes, result)
	}
	if relationships, hasRelationships := ro[memberRelationships]; hasRelationships {
		validateRelationshipsObject(relationships, result)
	}
	if links, hasLinks := ro[memberLinks]; hasLinks {
		validateLinksObject(links, result)
	}
	if meta, hasMeta := ro[memberMeta]; hasMeta {
		validateMetaObject(meta, result)
	}
}

func isResourceIdentifierObject(r map[string]interface{}) bool {
	for key := range r {
		switch key {
		case memberID, memberType, memberMeta:
			// do nothing
		default:
			return false
		}
	}
	return true
}

func validateResourceIdentifierObject(ro map[string]interface{}, result *Result) {
	if id, hasID := ro[memberID]; !hasID {
		result.AddError(ErrResourceObjectMissingID)
	} else {
		validateID(id, result)
	}

	if jType, hasType := ro[memberType]; !hasType {
		result.AddError(ErrResourceObjectMissingType)
	} else {
		validateType(jType, result)
	}

	if meta, hasMeta := ro[memberMeta]; hasMeta {
		validateMeta(meta, result)
	}

	return
}

func validateResourceLinkage(rl interface{}, result *Result) {
	if rl == nil {
		return // null for empty to-one relationships.
	}

	switch reflect.TypeOf(rl).Kind() {
	case reflect.Map:
		// Single
		r := rl.(map[string]interface{})
		validateResourceIdentifierObject(r, result)
	case reflect.Slice:
		// Array
		resources := rl.([]interface{})
		for _, resource := range resources {
			r := resource.(map[string]interface{})
			validateResourceIdentifierObject(r, result)
		}
	default:
		result.AddError(ErrInvalidResourceLinkage)
	}
}

func validateMeta(meta interface{}, result *Result) {
	return
}

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

func validateRelationshipsObject(r interface{}, result *Result) {
	// The value of the relationships key MUST be an object (a “relationships
	// object”).
	relationships, ok := r.(map[string]interface{})
	if !ok {
		result.AddError(ErrNotRelationshipsObject)
		return // cannot procceed
	}

	for _, relationshipObject := range relationships {
		// TODO: perhaps warn if pluaral && ![]
		validateRelationshipObject(relationshipObject, result)
	}
}

func validateRelationshipObject(ro interface{}, result *Result) {
	relationshipObject, ok := ro.(map[string]interface{})
	if !ok {
		result.AddError(ErrNotRelationshipObject)
		return // cannot procceed
	}

	links, hasLinks := relationshipObject[memberLinks]
	data, hasData := relationshipObject[memberData]
	meta, hasMeta := relationshipObject[memberMeta]

	if !hasLinks && !hasData && !hasMeta {
		result.AddError(ErrRelationshipsObjectOneRequiredMember)
		return // cannot procede
	}

	if hasLinks {
		validateLinksObject(links, result)
	}
	if hasData {
		validateResourceLinkage(data, result)
	}
	if hasMeta {
		validateMetaObject(meta, result)
	}
}

func validateLinksObject(l interface{}, result *Result) {
	fmt.Println("validateLinksObject")
}
func validateMetaObject(m interface{}, result *Result) {
	fmt.Println("validateMetaObject")
}
func validateJSONAPIObject(ja interface{}, result *Result) {
	if _, ok := ja.(map[string]interface{}); !ok {
		result.AddError(ErrNotJSONAPIObject)
	}
}

func validateErrors(errors map[string]interface{}, result *Result) {
	for k := range errors {
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
