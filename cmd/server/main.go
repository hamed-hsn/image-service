package main

import (
	"image_service/internal/adapter/db"
	"image_service/internal/app"
	"image_service/internal/controller/server"
	"image_service/internal/delivery/rest"
	"image_service/internal/usecase/metadata"
	"image_service/internal/validators"
	"image_service/pkg/simplelogger"
	"time"
)

func main() {
	logger := simplelogger.New()
	mongoDB := db.NewMongoDbAdapter(logger,
		db.WithURI(app.MongoDbUri),
		db.WithDbName(app.MongoDbDatabaseName),
		db.WithColName(app.MongoDbColName),
		db.WithPagination(10),
		db.WithCheckingDuraion(time.Minute*10),
	)
	metaUC := metadata.New()
	builder := server.NewBuilder()
	builder.Logger(logger).Db(mongoDB).MetaDataDetector(metaUC).
		UniqueValidator(validators.DefaultUniqueValidator).
		FilterValidator(validators.DefaultFilterValidator).
		UploadValidator(validators.DefaultUploadValidator)
	ctrl := builder.Build()
	e := rest.New(ctrl)
	err := e.Start(app.HttpPort)
	logger.Error("app crushed", "error", err)
}
