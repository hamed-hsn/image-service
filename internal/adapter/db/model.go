package db

import (
	"image_service/entity"
	"time"
)

type infoModel struct {
	Url          string         `json:"url" bson:"url"`
	LocalPath    string         `json:"local_path" bson:"local_path"`
	CommonKey    string         `json:"common_key" bson:"common_key"`
	Ext          string         `json:"ext" bson:"ext"`
	MimeType     string         `json:"mime-type" bson:"mime-type"`
	Size         uint64         `json:"size" bson:"size"`
	DownloadedAt uint64         `json:"downloaded_at" bson:"downloaded_at"`
	Mode         string         `json:"mode" bson:"mode"`
	Meta         map[string]any `json:"meta" bson:"meta"`
	DeletedAt    *time.Time     `json:"deleted_at" bson:"deleted_at"`
	UpdatedAt    *time.Time     `json:"updated_at" bson:"updated_at"`
}

func (i infoModel) toEntity() *entity.Info {
	return &entity.Info{
		Url:          i.Url,
		LocalPath:    i.LocalPath,
		CommonKey:    i.CommonKey,
		Ext:          i.Ext,
		MimeType:     i.MimeType,
		Size:         i.Size,
		DownloadedAt: i.DownloadedAt,
		Mode:         entity.ResolveMode(i.Mode),
		Meta:         i.Meta,
	}
}

func toModel(e *entity.Info) infoModel {
	return infoModel{
		Url:          e.Url,
		LocalPath:    e.LocalPath,
		CommonKey:    e.CommonKey,
		Ext:          e.Ext,
		MimeType:     e.MimeType,
		Size:         e.Size,
		DownloadedAt: e.DownloadedAt,
		Mode:         string(e.Mode),
		Meta:         e.Meta,
	}
}

func toEntities(models []infoModel) []*entity.Info {
	if len(models) == 0 {
		return nil
	}
	res := make([]*entity.Info, 0, len(models))
	for _, m := range models {
		res = append(res, m.toEntity())
	}
	return res
}
