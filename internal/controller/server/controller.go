package server

import (
	"context"
	"image_service/internal/dto"
	"image_service/internal/protocol"
)

func New(db protocol.Db, down protocol.DownloaderUC, meta protocol.MetadataUC, logger protocol.Logger) *Controller {
	return &Controller{
		db:         db,
		downloader: down,
		meta:       meta,
		logger:     logger,
	}
}

type Controller struct {
	db         protocol.Db
	downloader protocol.DownloaderUC
	meta       protocol.MetadataUC
	logger     protocol.Logger
}

func (c *Controller) ListImages(ctx context.Context, request dto.ListImageRequest) (dto.ListImageResponse, error) {
	panic("")
}

func (c *Controller) GetImageInfo(ctx context.Context, request dto.GetImageRequest) (dto.GetImageInfoResponse, error) {
	panic("")
}

func (c *Controller) ShowImage(ctx context.Context, request dto.GetImageRequest) (dto.GetImageRawResponse, error) {
	panic("")
}

func (c *Controller) UploadImage(ctx context.Context, request dto.UploadRequest) (dto.UploadResponse, error) {
	panic("")
}