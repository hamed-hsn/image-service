package downloader

import (
	"fmt"
	"time"
)

var (
	DefaultFetcher = simpleFetcher{
		timeout: 15 * time.Second,
	}

	userAgent = "Mozilla/5.0 (platform; rv:geckoversion) Gecko/geckotrail Firefox/firefoxversion"

	invalidWorkersMsg = fmt.Sprintf("the workers count must be greater than %d", MinimumWorkersCount)
)

const (
	DefaultInputChannelLen = 10
	MinimumWorkersCount    = 4
)
