//go:build go1.19 && !generate
// +build go1.19,!generate

package endian_test

import (
	"encoding/binary"
	"testing"

	"github.com/dolmen-go/endian"
)

// If endian.Native is not fully defined, this file should not compile.
// Check for support of all encoding/binary interfaces.
var compileTestAppend interface {
	binary.ByteOrder
	binary.AppendByteOrder
} = endian.Native

// TestCompile119 is a dummy test: the real test is above.
func TestCompileAppendByteOrder(t *testing.T) {
	_ = compileTest
}
