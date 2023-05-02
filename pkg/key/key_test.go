package key

import (
	"fmt"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	u := GenerateKey("hamed")
	fmt.Println(u)
}
