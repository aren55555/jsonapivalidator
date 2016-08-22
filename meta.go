package jsonapivalidator

func validateMetaObject(m interface{}, result *Result) {
	_, ok := m.(map[string]interface{})
	if !ok {
		result.AddError(ErrNotMetaObject)
		return // cannot procceed
	}
}
