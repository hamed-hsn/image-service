package dto

import (
	"image_service/entity"
	"io"
)

type ListImageRequest struct {
	Page   int   `json:"page" query:"page"`
	After  int64 `json:"after,omitempty" query:"after,omitempty"`
	Before int64 `json:"before,omitempty" query:"before,omitempty"`
}

type ListImageResponse struct {
	Data    []*entity.Info `json:"data"`
	Message string         `json:"message"`
	Source  string         `json:"source"`
}

type GetImageRequest struct {
	CommonKey string `json:"common_key,omitempty" query:"common_key,omitempty"`
	Url       string `json:"url,omitempty" query:"url,omitempty"`
}

type GetImageInfoResponse struct {
	Info    *entity.Info `json:"info,omitempty"`
	Message string       `json:"message"`
	Source  string       `json:"source"`
}

type GetImageRawResponse struct {
	Body     io.ReadCloser
	MimeType string
	Ext      string
}

type UploadRequest struct {
	File        io.ReadCloser
	ContentType string
	Size        int64
}

type UploadResponse struct {
	Info *entity.Info `json:"info,omitempty"`
}
