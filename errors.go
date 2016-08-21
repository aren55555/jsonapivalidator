package jsonapivalidator

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
