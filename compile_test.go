//go:build !generate
// +build !generate

package endian_test

import (
	"encoding/binary"
	"reflect"
	"testing"

	"github.com/dolmen-go/endian"
)

// If endian.Native is not defined, this file should not compile
var compileTest binary.ByteOrder = endian.Native

// TestCompile is a dummy test: the real test is above
func TestCompile(t *testing.T) {
	_ = compileTest
}

func TestEqual(t *testing.T) {
	nbOK := 0
	for _, bo := range []binary.ByteOrder{
		binary.BigEndian,
		binary.LittleEndian,
	} {
		t.Logf("Native == %v: %v", bo, endian.Native == bo)
		if endian.Native == bo {
			nbOK++
		}
	}
	if nbOK != 1 {
		t.Error("1 equal expected, got", nbOK)
	}
}

// go test -run TestDynamic -v -short
// go test -run TestDynamic -v -short -tags endian_nostatic
func TestDynamic(t *testing.T) {
	typ := reflect.TypeOf(&endian.Native).Elem()

	if typ.Kind() == reflect.Interface {
		// To test this case: go test -tags endian_nostatic -v -run TestDynamic
		t.Log("endian.Native is dynamic")
	} else {
		t.Log("endian.Native is static")
	}

	if testing.Short() {
		return
	}
	t.Log(typ)
}
