package server

import (
	"bytes"
	"context"
	"fmt"
	"image_service/entity"
	"image_service/internal/app"
	"image_service/internal/dto"
	"image_service/internal/protocol"
	"image_service/pkg/key"
	"io"
	"os"
	"time"
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
	if err := c.uploadValidator.Validate(request); err != nil {
		c.logger.Error("validate error", err)
		return dto.UploadResponse{}, err
	}
	raw, err := io.ReadAll(request.File)
	if err != nil {
		c.logger.Error("io read error", err)
		return dto.UploadResponse{}, err
	}
	buf := bytes.NewBuffer(raw)
	r, err := c.meta.DetectFromBuffer(buf)
	if err != nil {
		c.logger.Error("meta error", err)
		return dto.UploadResponse{}, err
	}

	info := entity.Info{
		DownloadedAt: uint64(time.Now().Unix()),
		CommonKey:    key.GenerateKey(""),
		Mode:         entity.DownloadedByUserRequestMode,
		Ext:          r.Ext,
		MimeType:     r.Mime,
		Size:         uint64(len(raw)),
	}
	path := fmt.Sprintf("%s/%s.%s", app.UploadedImagesDirPath, info.CommonKey, r.Ext)
	info.LocalPath = path
	if err = store(raw, path); err != nil {
		c.logger.Error("store error", err)
		return dto.UploadResponse{}, err
	}

	if err = c.db.Insert(ctx, &info); err != nil {
		c.logger.Error("db error", err)
		return dto.UploadResponse{}, err
	}
	return dto.UploadResponse{Info: &info}, nil
}

func store(raw []byte, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = f.Write(raw)
	return err
}
