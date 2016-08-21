package jsonapivalidator

func validateRelationshipsObject(r interface{}, result *Result) {
	// The value of the relationships key MUST be an object (a “relationships
	// object”).
	relationships, ok := r.(map[string]interface{})
	if !ok {
		result.AddError(ErrNotRelationshipsObject)
		return // cannot procceed
	}

	for _, relationshipObject := range relationships {
		// TODO: perhaps warn if pluaral && ![]
		validateRelationshipObject(relationshipObject, result)
	}
}

func validateRelationshipObject(ro interface{}, result *Result) {
	relationshipObject, ok := ro.(map[string]interface{})
	if !ok {
		result.AddError(ErrNotRelationshipObject)
		return // cannot procceed
	}

	links, hasLinks := relationshipObject[memberLinks]
	data, hasData := relationshipObject[memberData]
	meta, hasMeta := relationshipObject[memberMeta]

	if !hasLinks && !hasData && !hasMeta {
		result.AddError(ErrRelationshipsObjectOneRequiredMember)
		return // cannot procede
	}

	if hasLinks {
		validateLinksObject(links, result)
	}
	if hasData {
		validateResourceLinkage(data, result)
	}
	if hasMeta {
		validateMetaObject(meta, result)
	}
}
