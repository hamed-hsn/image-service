package validators

import "image_service/internal/protocol"

type uploadValidator struct {
}

var _ protocol.UploadValidator = uploadValidator{}
