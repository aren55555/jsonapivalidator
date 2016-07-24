package jsonapivalidator

import "reflect"

// Validate checks the payload against the JSONAPI sepc, returns a Result
// populated with all spec violations.
func Validate(payload interface{}) (result *Result) {
	result = NewResult()

	// Check the root
	root := payload.(map[string]interface{})
	data, dataExists := root["data"]
	errors, errorsExists := root["errors"]
	_, metaExists := root["meta"]

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
		case "id", "type", "attributes", "relationships", "links", "meta":
			// do nothing
		default:
			return false
		}
	}
	return true
}

func validateResourceObject(ro map[string]interface{}, result *Result) {
	if id, hasID := ro["id"]; !hasID {
		result.AddError(ErrResourceObjectMissingID)
	} else {
		validateID(id, result)
	}

	if jType, hasType := ro["type"]; !hasType {
		result.AddError(ErrResourceObjectMissingType)
	} else {
		validateType(jType, result)
	}

	// TODO: finish validate RO
}

func isResourceIdentifierObject(r map[string]interface{}) bool {
	for key := range r {
		switch key {
		case "id", "type", "meta":
			// do nothing
		default:
			return false
		}
	}
	return true
}

func validateResourceIdentifierObject(ro map[string]interface{}, result *Result) {
	// TODO: validate RIO
}

func validateID(id interface{}, result *Result) {
	// TODO: validate ID is a string
}

func validateType(t interface{}, result *Result) {
	// TODO: validate Type is a string
}

func validateErrors(errors map[string]interface{}, result *Result) {
	for k := range errors {
		switch k {
		case "id":
		case "links":
		case "status":
		case "code":
		case "title":
		case "detail":
		case "source":
		case "meta":
		default:
			result.AddError(ErrInvalidErrorMember)
		}
	}
}
