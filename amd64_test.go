//go:build amd64
// +build amd64

package endian_test

import (
	"encoding/binary"
	"testing"

	"github.com/dolmen-go/endian"
)

func TestAmd64(t *testing.T) {
	if endian.Native != binary.LittleEndian {
		t.Fatal("Unexpected native encoding:", endian.Native)
	}
}
