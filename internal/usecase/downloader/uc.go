package downloader

import (
	"context"
	"image_service/internal/protocol"
	"time"
)

func NewDefaultDownloaderUC(workers int, logger protocol.Logger) *concurrentDownloader {
	if workers < MinimumWorkersCount {
		panic(invalidWorkersMsg)
	}
	workersDone := make(map[int]chan struct{})
	for i := 0; i < workers; i++ {
		workersDone[i] = make(chan struct{})
	}
	return &concurrentDownloader{
		input:   make(chan string, DefaultInputChannelLen),
		errorC:  make(chan errorType, DefaultInputChannelLen),
		outputC: make(chan outputType, DefaultInputChannelLen),
		//done unbuffered channel for sync and flag
		done:            make(chan struct{}),
		logger:          logger,
		workers:         workers,
		fetcher:         DefaultFetcher,
		workersDoneFlag: workersDone,
	}
}

type concurrentDownloader struct {
	input           chan string
	errorC          chan errorType
	outputC         chan outputType
	done            chan struct{}
	workersDoneFlag map[int]chan struct{}
	logger          protocol.Logger
	workers         int
	fetcher         fetcher
}

func (cd *concurrentDownloader) Start() error {
	for i := 0; i < cd.workers; i++ {
		go func(workerID int) {
			for {
				select {
				case url := <-cd.input:
					cd.logger.Info("received new url to worker-id #", "worker-id", workerID)
					go cd.download(url)
				case <-cd.workersDoneFlag[workerID]:
					cd.logger.Info("worker received done flag", "worker-id", workerID)
					return
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
		cd.errorC <- errorType{Url: url, Status: status, Error: err}
		return
	}
	cd.outputC <- outputType{Body: body, Status: status, Url: url}
}

func (cd *concurrentDownloader) Input() chan string {
	return cd.input
}

func (cd *concurrentDownloader) Output() chan outputType {
	return cd.outputC
}

func (cd *concurrentDownloader) Errors() chan errorType {
	return cd.errorC
}

func (cd *concurrentDownloader) StopGracefully() {
	cd.logger.Info("stop gracefully start!")
	for _, flagChannel := range cd.workersDoneFlag {
		flagChannel <- struct{}{}
	}
	cd.done <- struct{}{}
	go func() {
		time.Sleep(time.Second * 5)
		close(cd.input)
		close(cd.outputC)
		close(cd.errorC)
	}()
}
