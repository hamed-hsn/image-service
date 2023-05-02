package types

type ListFilterDB struct {
	Page   *int64
	After  *uint64
	Before *uint64
}

type GetFilterDB struct {
	ByCommonKey *string
	ByUrl       *string
	ByPath      *string
}

type UpdateOneDB struct{}

type DeleteOneDB struct {
}
