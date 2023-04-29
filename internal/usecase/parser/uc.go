package parser

import (
	"bufio"
	"image_service/internal/protocol"
	"io"
	"sync"
)

type ioParser struct {
	reader    io.ReadCloser
	validator protocol.LinkValidator
	mu        sync.Mutex
	logger    protocol.Logger
}

func (i *ioParser) Parse() ([]string, error) {
	i.mu.Lock()
	defer i.mu.Unlock()
	defer func() {
		if err := i.reader.Close(); err != nil {
			i.logger.Error(err.Error())
		}
	}()
	scanner := bufio.NewScanner(i.reader)
	scanner.Split(bufio.ScanLines)
	urls := make([]string, 0, urlsDefaultSizse)
	for scanner.Scan() {
		if value := scanner.Text(); i.validator.ValidateLink(value) == nil {
			urls = append(urls, value)
		}
	}
	return urls, scanner.Err()
}

func NewIoParser(file io.ReadCloser, validator protocol.LinkValidator) *ioParser {
	return &ioParser{
		validator: validator,
		reader:    file,
	}
}
