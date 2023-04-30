package dto

import "io"

type DownloaderOutput struct {
	Body   io.ReadCloser
	Status int
	Url    string
}

type DownloaderError struct {
	Url    string
	Status int
	Error  error
}
