// +build mips64 ppc64 s390x

package endian

import "encoding/binary"

var Native = binary.BigEndian
