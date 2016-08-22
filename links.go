package jsonapivalidator

import "reflect"

func validateLinksObject(l interface{}, result *Result) {
	links, ok := l.(map[string]interface{})
	if !ok {
		result.AddError(ErrNotLinksObject)
		return // cannot procceed
	}

	for _, v := range links {
		validateLinkObject(v, result)
	}
}

func validateLinkObject(l interface{}, result *Result) {
	// Each member of a links object is a “link”. A link MUST be represented as
	// either:
	switch reflect.TypeOf(l).Kind() {
	case reflect.String:
		// a string containing the link’s URL.
		validateURL(l, result)
	case reflect.Map:
		// an object (“link object”) which can contain the following members:
		//  href: a string containing the link’s URL.
		//  meta: a meta object containing non-standard meta-information about the
		//  link.
		link, ok := l.(map[string]interface{})
		if !ok {
			result.AddError(ErrNotLinkObject)
			return // cannot proceed
		}

		for k, v := range link {
			switch k {
			case memberHref:
				// a string containing the link’s URL.
				validateURL(v, result)
			case memberMeta:
				// a meta object containing non-standard meta-information about the
				// link.
				validateMetaObject(v, result)
			default:
				result.AddError(ErrInvalidLinkMember)
			}
		}
	default:
		result.AddError(ErrInvalidLinkType)
	}
}
