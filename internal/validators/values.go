package validators

import "errors"

var (
	ErrLessThanMinimumLength = errors.New("minimum length control failed")
	ErrInvalidScheme         = errors.New("invalid scheme")
)

var (
	HttpsScheme = "https://"
	HttpScheme  = "http://"
)
