package validators

import (
	"image_service/internal/dto"
	"image_service/internal/protocol"
)

type filterValidator struct {
}

func (f filterValidator) ValidateListReq(request dto.ListImageRequest) error {
	if request.Page < 0 {
		return ErrInvalidFilter
	}
	return nil
}

func (f filterValidator) ValidateGetInfoReq(request dto.GetImageRequest) error {
	if request.Url == "" && request.CommonKey == "" {
		return ErrInvalidFilter
	}
	return nil
}

var _ protocol.FilterValidator = filterValidator{}
