package jsonapivalidator

import "reflect"

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
		result.AddError(ErrResourceObjectIdentifierMissingID)
	} else {
		validateID(id, result)
	}

	if jType, hasType := ro[memberType]; !hasType {
		result.AddError(ErrResourceObjectIdentifierMissingType)
	} else {
		validateType(jType, result)
	}

	if meta, hasMeta := ro[memberMeta]; hasMeta {
		validateMetaObject(meta, result)
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
