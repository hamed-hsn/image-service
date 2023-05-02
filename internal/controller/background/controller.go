package background

import (
	"bytes"
	"context"
	"fmt"
	"image_service/entity"
	"image_service/internal/app"
	"image_service/internal/dto"
	"image_service/internal/protocol"
	"image_service/internal/types"
	"image_service/pkg/key"
	"io"
	"os"
	"time"
)

func New(db protocol.Db, downloader protocol.DownloaderUC, meta protocol.MetadataUC, parser protocol.ParserUC, logger protocol.Logger) *BackgroundCtrl {
	return &BackgroundCtrl{
		db:         db,
		downloader: downloader,
		metadata:   meta,
		parser:     parser,
		done:       make(chan struct{}),
		logger:     logger,
	}
}

type BackgroundCtrl struct {
	db         protocol.Db
	downloader protocol.DownloaderUC
	metadata   protocol.MetadataUC
	parser     protocol.ParserUC
	done       chan struct{}
	logger     protocol.Logger
}

func (b *BackgroundCtrl) Feed(ctx context.Context) {
	links, err := b.parser.Parse()
	if err != nil {
		b.logger.Error("bg-controller, parse error", "error", err)
		return
	}
	for _, link := range links {
		if info, err := b.db.Get(ctx, types.GetFilterDB{ByUrl: &link}); err == nil {
			b.logger.Warning("already info exist by this url", "url", link, "common-key", info.CommonKey)
			continue
		}
		b.downloader.Input() <- link
	}
}

func (b *BackgroundCtrl) LinkProcess(out *dto.DownloaderOutput) (*entity.Info, error) {
	info := entity.Info{
		Url:          out.Url,
		DownloadedAt: uint64(time.Now().Unix()),
		CommonKey:    key.GenerateKey(out.Url),
	}
	defer out.Body.Close()
	raw, err := io.ReadAll(out.Body)
	if err != nil {
		b.logger.Error("can not read from buffer", "error", err)
		return nil, err
	}
	info.Size = uint64(len(raw))
	buf := bytes.NewBuffer(raw)
	meta, err := b.metadata.DetectFromBuffer(buf)
	if err != nil {
		b.logger.Error("metadata error", "error", err)
	}
	info.MimeType = meta.Mime
	info.Ext = meta.Ext
	path := fmt.Sprintf("%s/%s.%s", app.ImagesFsPath, info.CommonKey, meta.Ext)
	info.LocalPath = path
	if err := store(raw, path); err != nil {
		b.logger.Error("bg-controller - store - os.Create", "error", err)
		return nil, err
	}
	info.Mode = entity.DownloadedFromImageListFileMode
	return &info, nil
}

func store(raw []byte, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = f.Write(raw)
	return err
}

func (b *BackgroundCtrl) Process(ctx context.Context) {
	for {
		select {
		case out := <-b.downloader.Output():
			if info, err := b.LinkProcess(out); err == nil {
				err = b.db.Insert(ctx, info)
				if err != nil {
					b.logger.Error("inserting db err", "err", err)
				}
			}
		case err := <-b.downloader.Errors():
			b.logger.Error("bg error", "error", err.Error, "url", err.Url)
		}
	}
}

func (b *BackgroundCtrl) Start() {
	b.logger.Info("start background")
	ctx := context.Background()
	go b.downloader.Start()
	b.Feed(ctx)
	go b.Process(ctx)
	<-b.done
	b.downloader.StopGracefully()
	b.logger.Info("end of background")
}

func (b *BackgroundCtrl) Stop() {
	b.done <- struct{}{}
}
