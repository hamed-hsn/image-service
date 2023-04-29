package validators

import (
	"image_service/internal/protocol"
	"strings"
)

type linkValidator struct {
	minLength int
}

func (l linkValidator) ValidateLink(s string) error {
	s = strings.TrimSpace(s)
	if len(s) <= l.minLength {
		return ErrLessThanMinimumLength
	}
	if strings.HasPrefix(s, HttpsScheme) || strings.HasPrefix(s, HttpScheme) {
		return nil
	}
	return ErrInvalidScheme
}

var _ protocol.LinkValidator = linkValidator{}
