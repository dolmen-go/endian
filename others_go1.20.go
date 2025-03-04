//go:build go1.20 && !generate && !endian_nostatic
// +build go1.20,!generate,!endian_nostatic

package endian

import "encoding/binary"

// Native is the byte order of GOARCH.
var Native = binary.NativeEndian
