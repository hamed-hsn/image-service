package server

import (
	"errors"
	"image_service/internal/protocol"
)

func NewBuilder() *builder {
	return &builder{}
}

type builder struct {
	logger          protocol.Logger
	db              protocol.Db
	meta            protocol.MetadataUC
	uniqueValidator protocol.UniqueValidator
	filterValidator protocol.FilterValidator
	uploadValidator protocol.UploadValidator
}

func (b *builder) Logger(l protocol.Logger) *builder {
	if l == nil {
		panic("invalid logger")
	}
	b.logger = l
	return b
}

func (b *builder) Db(db protocol.Db) *builder {
	if db == nil {
		panic("invalid db")
	}
	b.db = db
	return b
}

func (b *builder) MetaDataDetector(m protocol.MetadataUC) *builder {
	if m == nil {
		panic("invalid metadata")
	}
	b.meta = m
	return b
}

func (b *builder) UniqueValidator(v protocol.UniqueValidator) *builder {
	if v == nil {
		panic("invalid unique validator")
	}
	b.uniqueValidator = v
	return b
}
func (b *builder) FilterValidator(v protocol.FilterValidator) *builder {
	if v == nil {
		panic("invalid filter validator")
	}
	b.filterValidator = v
	return b
}

func (b *builder) UploadValidator(v protocol.UploadValidator) *builder {
	if v == nil {
		panic("invalid upload validator")
	}
	b.uploadValidator = v
	return b
}

func (b *builder) check() error {
	if b.logger == nil {
		return errors.New("logger is nil")
	}
	if b.meta == nil {
		return errors.New("meta is nil")
	}
	if b.db == nil {
		return errors.New("db is nil")
	}
	if b.uniqueValidator == nil {
		return errors.New("unique validator is nil")
	}
	if b.filterValidator == nil {
		return errors.New("filter validator is nil")
	}
	if b.uploadValidator == nil {
		return errors.New("upload validator is nil")
	}
	return nil
}

func (b *builder) Build() *Controller {
	if err := b.check(); err != nil {
		panic(err)
	}
	return &Controller{
		db:              b.db,
		meta:            b.meta,
		uniqueValidator: b.uniqueValidator,
		filterValidator: b.filterValidator,
		uploadValidator: b.uploadValidator,
		logger:          b.logger,
	}
}
