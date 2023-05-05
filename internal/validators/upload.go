package validators

import (
	"image_service/internal/dto"
	"image_service/internal/protocol"
)

type uploadValidator struct {
	MaxSize             int64
	InvalidContentTypes []string
}

func (u uploadValidator) Validate(request dto.UploadRequest) error {
	if request.Size > u.MaxSize {
		return ErrInvalidSize
	}
	for _, mt := range u.InvalidContentTypes {
		if mt == request.ContentType {
			return ErrInvalidContentTypes
		}
	}
	return nil
}

var _ protocol.UploadValidator = uploadValidator{}
