//go:build amd64 && !generate
// +build amd64,!generate

package endian_test

import (
	"encoding/binary"
	"fmt"
	"runtime"

	"github.com/dolmen-go/endian"
)

func Example_little() {
	fmt.Printf("GOARCH=%s: %s\n", runtime.GOARCH, endian.Native)

	const n = 0xDeadBeef
	var b [4]byte
	endian.Native.PutUint32(b[:], n)
	fmt.Printf("0x%x => [% x]\n", n, b)
	// Output:
	// GOARCH=amd64: LittleEndian
	// 0xdeadbeef => [ef be ad de]
}

func Example_littleEqual() {
	for _, bo := range []binary.ByteOrder{
		binary.BigEndian,
		binary.LittleEndian,
	} {
		fmt.Println("endian.Native ==", bo, "=>", endian.Native == bo)
	}
	// Output:
	// endian.Native == BigEndian => false
	// endian.Native == LittleEndian => true
}
