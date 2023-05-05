package validators

import "errors"

var (
	ErrLessThanMinimumLength = errors.New("minimum length control failed")
	ErrInvalidScheme         = errors.New("invalid scheme")

	ErrInvalidFilter       = errors.New("invalid filter")
	ErrInvalidSize         = errors.New("invalid size")
	ErrInvalidContentTypes = errors.New("content-type is forbidden")
)

var (
	HttpsScheme = "https://"
	HttpScheme  = "http://"
)
