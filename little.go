// +build 386 amd64 amd64p32 arm arm64 mips64le ppc64le

package endian

import "encoding/binary"

var Native = binary.LittleEndian