package protocol

import "image_service/internal/dto"

type LinkValidator interface {
	ValidateLink(string) error
}

type UniqueValidator interface {
}

type FilterValidator interface {
	ValidateListReq(request dto.ListImageRequest) error
	ValidateGetInfoReq(request dto.GetImageRequest) error
}

type UploadValidator interface {
	Validate(request dto.UploadRequest) error
}
