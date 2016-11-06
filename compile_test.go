package endian_test

import (
	"encoding/binary"
	"testing"

	"github.com/dolmen-go/endian"
)

// If endian.Native is not defined, this file should not compile
var compileTest binary.ByteOrder = endian.Native

// TestCompile is a dummy test: the real test is above
func TestCompile(t *testing.T) {
}
