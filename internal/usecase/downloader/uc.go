package downloader

import (
	"context"
	"image_service/internal/protocol"
	"io"
)

type concurrentDownloader struct {
	input  chan string
	ErrorC chan struct {
		Url    string
		Status int
		Error  error
	}
	Output chan struct {
		Body   io.ReadCloser
		Status int
	}
	done    chan struct{}
	logger  protocol.Logger
	workers int
	fetcher fetcher
}

func (cd *concurrentDownloader) Start() error {
	for i := 0; i < cd.workers; i++ {
		go func(workerID int) {
			for {
				select {
				case url := <-cd.input:
					cd.logger.Info("received new url to worker-id #", "worker-id", workerID)
					go cd.download(url)
				}
			}
		}(i)
	}
	<-cd.done
	return nil
}

func (cd *concurrentDownloader) download(url string) {
	ctx := context.Background()
	body, status, err := cd.fetcher.fetch(ctx, url)
	if err != nil || status != 200 {
		cd.ErrorC <- struct {
			Url    string
			Status int
			Error  error
		}{Url: url, Status: status, Error: err}
		return
	}
	cd.Output <- struct {
		Body   io.ReadCloser
		Status int
	}{Body: body, Status: status}
}
