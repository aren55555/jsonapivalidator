package jsonapivalidator

import "reflect"

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
