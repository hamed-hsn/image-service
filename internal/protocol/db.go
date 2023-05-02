package protocol

import (
	"context"
	"image_service/entity"
	"image_service/internal/types"
)

type Db interface {
	List(ctx context.Context, filter types.ListFilterDB) ([]*entity.Info, error)
	Get(ctx context.Context, filter types.GetFilterDB) (*entity.Info, error)
	Insert(ctx context.Context, info *entity.Info) error
	UpdateOne(ctx context.Context, update types.UpdateOneDB) (*entity.Info, error)
	DeleteOne(ctx context.Context, filter types.DeleteOneDB) (entity.Info, error)
}
