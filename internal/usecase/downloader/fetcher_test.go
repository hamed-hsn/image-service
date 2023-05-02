package downloader

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestSimpleFetcher(t *testing.T) {
	url := "https://tourism.780.ir/tourism/_next/static/media/domestic-landing.d33edef1.png"
	body, status, err := DefaultFetcher.fetch(context.TODO(), url)
	assert.Nil(t, err)
	assert.Equal(t, status, 200)
	raw, _ := io.ReadAll(body)
	fmt.Println(string(raw))
}
