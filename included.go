package jsonapivalidator

import "fmt"

func validateIncluded(included interface{}, result *Result) {
	// included MUST be an array
	includes, ok := included.([]interface{})
	if !ok {
		result.AddError(ErrInvalidIncludedType)
	}

	for _, i := range includes {
		fmt.Println(i)
	}
}
