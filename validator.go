package jsonapivalidator

import "reflect"

// Validate checks the payload against the JSONAPI sepc, returns a Result
// populated with all spec violations.
func Validate(payload interface{}) (result *Result) {
	result = NewResult()

	// Check the root
	root := payload.(map[string]interface{})
	data, dataExists := root[memberData]
	errors, errorsExists := root[memberErrors]
	_, metaExists := root[memberMeta]

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
	return
}

func validateData(data interface{}, result *Result) {
	if data == nil {
		return
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

	// TODO: finish validate RO
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
