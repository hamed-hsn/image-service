package parser

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"image_service/internal/protocol"
	"io"
	"testing"
)

func TestParse(t *testing.T) {
	reader := bytes.NewReader(fakeContent)
	p := NewIoParser(io.NopCloser(reader), mockValidator{}, nil)
	links, err := p.Parse()
	assert.Nil(t, err)
	assert.Equal(t, len(links), 5)
	assert.Contains(t, links, "123")
}

type mockValidator struct {
}

func (m mockValidator) ValidateLink(s string) error {
	if s == "this-is-not-valid" {
		return errors.New("fake error")
	}
	return nil
}

var _ protocol.LinkValidator = mockValidator{}

var fakeContent = []byte(`123
aaa
this-is-not-valid
ccc
aaa
ddd
`)
