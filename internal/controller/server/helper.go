package server

import (
	"image_service/internal/dto"
	"image_service/internal/types"
)

func makeListFilter(request dto.ListImageRequest) types.ListFilterDB {
	filter := types.ListFilterDB{}
	if request.Page > 0 {
		i := int64(request.Page)
		filter.Page = &i
	}
	if request.After > 0 {
		i := uint64(request.After)
		filter.After = &i
	}
	if request.Before > 0 {
		i := uint64(request.Before)
		filter.Before = &i
	}
	return filter
}

func makeGetFilter(request dto.GetImageRequest) types.GetFilterDB {
	filter := types.GetFilterDB{}
	if request.Url != "" {
		u := request.Url
		filter.ByUrl = &u
	}
	if request.CommonKey != "" {
		u := request.CommonKey
		filter.ByCommonKey = &u
	}
	return filter
}
