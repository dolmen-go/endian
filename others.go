// Code generated by generate_go1.18.go; DO NOT EDIT.

//go:build !386 && !amd64 && !amd64p32 && !arm && !arm64 && !arm64be && !armbe && !loong64 && !mips && !mips64 && !mips64le && !mips64p32 && !mips64p32le && !mipsle && !ppc && !ppc64 && !ppc64le && !riscv && !riscv64 && !s390 && !s390x && !sparc && !sparc64 && !wasm && !generate
// +build !386,!amd64,!amd64p32,!arm,!arm64,!arm64be,!armbe,!loong64,!mips,!mips64,!mips64le,!mips64p32,!mips64p32le,!mipsle,!ppc,!ppc64,!ppc64le,!riscv,!riscv64,!s390,!s390x,!sparc,!sparc64,!wasm,!generate

package endian

import (
	"encoding/binary"
	"unsafe"
)

// Native is the byte order of GOARCH.
// It will be determined at runtime because it was unknown at code
// generation time.
var Native binary.ByteOrder

func init() {
	// http://grokbase.com/t/gg/golang-nuts/129jhmdb3d/go-nuts-how-to-tell-endian-ness-of-machine#20120918nttlyywfpl7ughnsys6pm4pgpe
	var x int32 = 0x01020304
	switch *(*byte)(unsafe.Pointer(&x)) {
	case 1:
		Native = binary.BigEndian
	case 4:
		Native = binary.LittleEndian
	}
}
