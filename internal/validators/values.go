package validators

import "errors"

var (
	ErrLessThanMinimumLength = errors.New("minimum length control failed")
	ErrInvalidScheme         = errors.New("invalid scheme")

	ErrInvalidFilter = errors.New("invalid filter")
)

var (
	HttpsScheme = "https://"
	HttpScheme  = "http://"
)
