package main

import (
	"image_service/internal/adapter/db"
	"image_service/internal/app"
	"image_service/internal/controller/background"
	"image_service/internal/usecase/downloader"
	"image_service/internal/usecase/metadata"
	"image_service/internal/usecase/parser"
	"image_service/internal/validators"
	"image_service/pkg/simplelogger"
	"os"
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
	file, err := os.Open(app.ImagesFilePath)
	if err != nil {
		logger.Error("can not open file", "error", err)
	}
	p := parser.NewIoParser(file, validators.DefaultLinkValidator, logger)
	downloaderUC := downloader.NewDefaultDownloaderUC(app.WorkersCount, logger)
	metaUC := metadata.New()
	ctrl := background.New(mongoDB, downloaderUC, metaUC, p, logger)

	ctrl.Start()
	defer ctrl.Stop()

}
