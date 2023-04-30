package downloader

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"image_service/internal/protocol"
	"io"
	"sync"
	"testing"
)

func newFakeConcurrentDownloader() *concurrentDownloader {
	wdf := make(map[int]chan struct{})
	wdf[0] = make(chan struct{})
	wdf[1] = make(chan struct{})
	return &concurrentDownloader{
		input:           make(chan string, 4),
		outputC:         make(chan outputType, 4),
		errorC:          make(chan errorType, 4),
		done:            make(chan struct{}),
		fetcher:         &mockFetcher{},
		workers:         4,
		logger:          &mockLogger{lastMetaData: make([]any, 0, 5)},
		workersDoneFlag: wdf,
	}
}

func TestConcurrentDownloader(t *testing.T) {
	cd := newFakeConcurrentDownloader()
	urls := []string{"some-image.com", "bad-url", "another-image.com",
		"some-image.com1", "bad-url", "another-image.com1",
		"some-image.com2", "bad-url", "another-image.com2",
		"some-image.com3", "bad-url", "another-image.com3",
		"some-image.com4", "bad-url", "another-image.com4",
		"some-image.com5", "bad-url", "another-image.com5",
		"some-image.com6", "bad-url", "another-image.com6"}
	go func() {
		err := cd.Start()
		assert.Nil(t, err)
	}()

	wg := sync.WaitGroup{}
	wg.Add(len(urls))
	fmt.Println("start")
	go func() {
		for _, url := range urls {
			cd.Input() <- url
		}
	}()
	go func() {
		for {
			select {
			case out := <-cd.Output():
				b, _ := io.ReadAll(out.Body)
				fmt.Println(string(b), out.Status, out.Url)
				_ = out.Body.Close()
				wg.Done()
			case err := <-cd.Errors():
				fmt.Println(err.Error, err.Url, err.Status)
				wg.Done()
			}
		}
	}()
	cd.StopGracefully()
	wg.Wait()
}

type mockLogger struct {
	lastMessage  string
	lastMetaData []any
}

var _ protocol.Logger = &mockLogger{}

func (m *mockLogger) Info(s string, a ...any) {
	fmt.Print("in mock info ", s, "  args: ")
	fmt.Println(a...)
	m.lastMessage = s
	m.lastMetaData = a
}

func (m *mockLogger) Warning(s string, a ...any) {
	//TODO implement me
	panic("implement me")
}

func (m *mockLogger) Debug(s string, a ...any) {
	//TODO implement me
	panic("implement me")
}

func (m *mockLogger) Error(s string, a ...any) {
	fmt.Print("in mock error ", s, "  args: ")
	fmt.Println(a...)
	m.lastMessage = s
	m.lastMetaData = a
}
