// +build !386,!amd64,!amd64p32,!arm,!arm64,!mips64,!mips64le,!ppc64,!ppc64le,!s390x,!others

package endian

import (
	"encoding/binary"
	"unsafe"
)

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