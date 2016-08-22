package jsonapivalidator

const (
	memberJSONAPI       = "jsonapi"
	memberData          = "data"
	memberErrors        = "errors"
	memberID            = "id"
	memberType          = "type"
	memberAttributes    = "attributes"
	memberRelationships = "relationships"
	memberLinks         = "links"
	memberMeta          = "meta"
	memberIncluded      = "included"

	// Error object specific
	memberStatus = "status"
	memberCode   = "code"
	memberTitle  = "title"
	memberDetail = "detail"
	memberSource = "source"

	// Links object specific
	memberSelf            = "self"
	memberRelated         = "related"
	memberPaginationFirst = "first"
	memberPaginationLast  = "last"
	memberPaginationPrev  = "prev"
	memberPaginationNext  = "next"

	// Link object specific
	memberHref = "href"
)
