package validators

import "image_service/internal/protocol"

type uniqueValidator struct{}

var _ protocol.UniqueValidator = uniqueValidator{}
