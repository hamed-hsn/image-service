package downloader

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
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

func (s *simpleFetcher) fetch(ctx context.Context, url string) (io.ReadCloser, int, error) {
	client := http.Client{}
	if s.timeout.Seconds() > 3 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, s.timeout)
		defer cancel()
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, 0, err
	}
	request.Header.Set("User-Agent", userAgent)
	resp, err := client.Do(request)
	if err != nil {
		return nil, 0, err
	}
	raw, err := io.ReadAll(resp.Body)
	buf := bytes.NewBuffer(raw)
	return io.NopCloser(buf), resp.StatusCode, err
}

var _ fetcher = &simpleFetcher{}

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
