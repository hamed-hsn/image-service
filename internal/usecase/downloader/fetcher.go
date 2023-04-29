package downloader

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"
)

type fetcher interface {
	fetch(ctx context.Context, url string) (io.ReadCloser, int, error)
}

type mockFetcher struct{}

func (m mockFetcher) fetch(_ context.Context, url string) (io.ReadCloser, int, error) {
	if url == "bad-url" {
		return nil, 0, fmt.Errorf("fetching url failed")
	}
	if url == "404" {
		return nil, 404, fmt.Errorf("not found")
	}
	body := bytes.NewReader([]byte("downloaded content"))
	return io.NopCloser(body), 200, nil
}

var _ fetcher = mockFetcher{}

type simpleFetcher struct {
	timeout time.Duration
}

func (s simpleFetcher) fetch(ctx context.Context, url string) (io.ReadCloser, int, error) {
	//TODO implement me
	panic("implement me")
}

var _ fetcher = simpleFetcher{}

type proxyFetcher struct {
	proxy struct {
		Http  string `json:"http,omitempty"`
		Https string `json:"https"`
	}
	timeout time.Duration
}

func (p proxyFetcher) fetch(ctx context.Context, url string) (io.ReadCloser, int, error) {
	//TODO implement me
	panic("implement me")
}

var _ fetcher = proxyFetcher{}
