package metadata

import (
	"bytes"
	"github.com/gabriel-vasile/mimetype"
	"image_service/internal/dto"
	"io"
	"strings"
	"sync"
)

type service struct {
	mu sync.Mutex
}

func (s *service) DetectFromBuffer(buff *bytes.Buffer) (dto.Metadata, error) {

	body, err := io.ReadAll(buff)
	if err != nil {
		return dto.Metadata{}, err
	}
	detector := mimetype.Detect(body)
	return dto.Metadata{
		Mime: detector.String(),
		Ext:  strings.ReplaceAll(detector.Extension(), ".", ""),
	}, nil
}

func New() *service {
	return &service{}
}
