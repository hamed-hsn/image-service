package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"image_service/internal/types"
)

func makeGetFilter(filter types.GetFilterDB) bson.M {
	f := bson.M{}
	if filter.ByUrl != nil {
		f["url"] = *filter.ByUrl
	}
	if filter.ByCommonKey != nil {
		f["common_key"] = *filter.ByCommonKey
	}
	if filter.ByPath != nil {
		f["local_path"] = *filter.ByPath
	}
	f["deleted_at"] = nil
	return f
}

func makeListFilter(filter types.ListFilterDB, pagination int64) (bson.M, *options.FindOptions) {
	opts := options.Find()
	conds := make([]bson.M, 0, 2)
	opts.SetSort(bson.M{"$natural": -1})
	page := int64(1)
	if filter.Page != nil {
		page = *filter.Page
	}
	opts.SetSkip((page - 1) * pagination)
	opts.SetLimit(pagination)
	if filter.After != nil {
		conds = append(conds, bson.M{"downloaded_at": bson.M{"$gt": *filter.After}})
	}
	if filter.Before != nil {
		conds = append(conds, bson.M{"downloaded_at": bson.M{"$lte": *filter.Before}})
	}
	conds = append(conds, bson.M{"deleted_at": nil})
	if len(conds) == 1 {
		return bson.M{"deleted_at": nil}, opts
	}
	return bson.M{"$and": conds}, opts
}
