package db

import (
	"context"
	"image_service/entity"
	"image_service/internal/types"
)

func (a *adapter) List(ctx context.Context, filter types.ListFilterDB) ([]*entity.Info, error) {
	bsonFilter, opts := makeListFilter(filter, a.pagination)
	result, err := a.getCol().Find(ctx, bsonFilter, opts)
	if err != nil {
		a.logger.Error("list-db", "error", err)
		return nil, err
	}
	var models []infoModel
	if err := result.All(ctx, &models); err != nil {
		a.logger.Error("list-db on decoding", "error", err)
		return nil, err
	}
	return toEntities(models), nil
}

func (a *adapter) Get(ctx context.Context, filter types.GetFilterDB) (*entity.Info, error) {
	bsonFilter := makeGetFilter(filter)
	result := a.getCol().FindOne(ctx, bsonFilter)
	if result.Err() != nil {
		//a.logger.Error("db-get error", "error", result.Err())
		return nil, result.Err()
	}
	var model infoModel
	if err := result.Decode(&model); err != nil {
		a.logger.Error("db-get error on decode data", "error", err)
		return nil, err
	}
	return model.toEntity(), nil
}

func (a *adapter) Insert(ctx context.Context, info *entity.Info) error {
	model := toModel(info)
	_, err := a.getCol().InsertOne(ctx, model)
	return err
}

func (a *adapter) UpdateOne(ctx context.Context, update types.UpdateOneDB) (*entity.Info, error) {
	panic("not implemented yet")
}

func (a *adapter) DeleteOne(ctx context.Context, filter types.DeleteOneDB) (entity.Info, error) {
	panic("not implemented yet")
}
