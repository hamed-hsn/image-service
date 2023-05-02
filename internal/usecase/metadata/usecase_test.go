package metadata

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestDetector(t *testing.T) {
	// the following path  is set to gitignore !!
	raw, err := os.Open("../../../files/tests/file.jpg")
	assert.Nil(t, err)
	d := New()
	buff := bytes.Buffer{}
	io.Copy(&buff, raw)
	ret, err := d.DetectFromBuffer(&buff)
	assert.Nil(t, err)
	assert.Equal(t, ret.Mime, "image/jpeg")
	assert.Equal(t, ret.Ext, "jpg")
}
