package jsonapivalidator

import (
	"fmt"
	"reflect"
)

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
		ro := data.(map[string]interface{})
		validateResourceObject(ro, result)
	case reflect.Slice:
		for _, ro := range data.([]map[string]interface{}) {
			validateResourceObject(ro, result)
		}
	default:
		result.AddError(ErrInvalidDataType)
	}
}

func validateResourceObject(ro map[string]interface{}, result *Result) {
	fmt.Println(ro)
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
