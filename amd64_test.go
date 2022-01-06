//go:build amd64 && !generate
// +build amd64,!generate

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
