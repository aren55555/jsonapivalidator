package jsonapivalidator

import "errors"

var (
	// ErrAtLeastOneRoot is the error returned when a document is missing a top level member
	ErrAtLeastOneRoot = errors.New("A document MUST contain at least one of the following top-level members: /data, /errors or /meta")
	// ErrRootDataAndErrors adsfasd adsf
	ErrRootDataAndErrors = errors.New("The members /data and /errors MUST NOT coexist in the same document")
	// ErrInvalidErrorMember is for when the errors object has an unexpected member key
	ErrInvalidErrorMember = errors.New("Invalid member to /errors")
	// ErrInvalidDataType is for data not being a hash array or null JSON type
	ErrInvalidDataType = errors.New("/data must contain a value that is a {}, [] or null")

	// ErrResourceObjectMissingID is returned when a resouce object did not have any value for the id key
	ErrResourceObjectMissingID = errors.New("A resource object MUST contain an id")
	// ErrResourceObjectMissingType is returned when a resouce object did not have any value for the type key
	ErrResourceObjectMissingType = errors.New("A resource object MUST contain a type")
	ErrNotAResource              = errors.New("Was not a resource object or a single resource identifier object")
)
