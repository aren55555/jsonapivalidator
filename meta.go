package jsonapivalidator

import "fmt"

func validateMetaObject(m interface{}, result *Result) {
	fmt.Println("validateMetaObject")
	_, ok := m.(map[string]interface{})
	if !ok {
		result.AddError(ErrNotMetaObject)
		return // cannot procceed
	}
}
