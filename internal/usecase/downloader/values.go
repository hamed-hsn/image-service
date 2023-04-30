package downloader

import (
	"fmt"
	"io"
	"time"
)

var (
	DefaultFetcher = simpleFetcher{
		timeout: 15 * time.Second,
	}

	invalidWorkersMsg = fmt.Sprintf("the workers count must be greater than %d", MinimumWorkersCount)
)

const (
	DefaultInputChannelLen = 10
	MinimumWorkersCount    = 4
)

type errorType struct {
	Url    string
	Status int
	Error  error
}

type outputType struct {
	Body   io.ReadCloser
	Status int
}
