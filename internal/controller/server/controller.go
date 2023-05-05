package server

import (
	"context"
	"image_service/internal/dto"
	"image_service/internal/protocol"
	"os"
)

func New(db protocol.Db, meta protocol.MetadataUC, logger protocol.Logger) *Controller {
	return &Controller{
		db:     db,
		meta:   meta,
		logger: logger,
	}
}

type Controller struct {
	db              protocol.Db
	meta            protocol.MetadataUC
	uniqueValidator protocol.UniqueValidator
	filterValidator protocol.FilterValidator
	uploadValidator protocol.UploadValidator
	logger          protocol.Logger
}

func (c *Controller) ListImages(ctx context.Context, request dto.ListImageRequest) (dto.ListImageResponse, error) {
	//	validate request
	if err := c.filterValidator.ValidateListReq(request); err != nil {
		c.logger.Error("in controller - validate error", "error", err)
		return dto.ListImageResponse{}, err
	}
	filter := makeListFilter(request)
	data, err := c.db.List(ctx, filter)
	if err != nil {
		c.logger.Error("in controller - list from db", "error", err)
		return dto.ListImageResponse{}, err
	}
	return dto.ListImageResponse{Data: data, Message: SuccessMessage, Source: DbSource}, nil
}

func (c *Controller) GetImageInfo(ctx context.Context, request dto.GetImageRequest) (dto.GetImageInfoResponse, error) {
	//	validate
	if err := c.filterValidator.ValidateGetInfoReq(request); err != nil {
		c.logger.Error("in controller - validate error", "error", err)
		return dto.GetImageInfoResponse{}, err
	}
	filter := makeGetFilter(request)
	info, err := c.db.Get(ctx, filter)
	if err != nil {
		c.logger.Error("in controller - db error")
		return dto.GetImageInfoResponse{}, err
	}
	return dto.GetImageInfoResponse{Info: info, Message: SuccessMessage, Source: DbSource}, nil
}

func (c *Controller) ShowImage(ctx context.Context, request dto.GetImageRequest) (dto.GetImageRawResponse, error) {
	if err := c.filterValidator.ValidateGetInfoReq(request); err != nil {
		c.logger.Error("in controller - validate error", "error", err)
		return dto.GetImageRawResponse{}, err
	}
	filter := makeGetFilter(request)
	info, err := c.db.Get(ctx, filter)
	if err != nil || info == nil {
		c.logger.Error("in controller - db error")
		return dto.GetImageRawResponse{}, err
	}
	file, err := os.Open(info.LocalPath)
	if err != nil {
		c.logger.Error("in controller / os error", "error", err)
		return dto.GetImageRawResponse{}, err
	}
	return dto.GetImageRawResponse{
		Body:     file,
		MimeType: info.MimeType,
		Ext:      info.Ext,
	}, nil
}

func (c *Controller) UploadImage(ctx context.Context, request dto.UploadRequest) (dto.UploadResponse, error) {
	panic("")
}
