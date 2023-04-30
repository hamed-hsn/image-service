package protocol

import (
	"image_service/internal/dto"
)

type ParserUC interface {
	Parse() ([]string, error)
}

type DownloaderUC interface {
	Start() error
	Input() chan string
	Output() chan dto.DownloaderOutput
	Errors() chan dto.DownloaderError
}
