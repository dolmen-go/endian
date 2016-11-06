//go:generate go run -tags generate generate.go

// Package endian provides the native byte order of GOARCH
//
// You should not use this package before reading and understanding
// https://commandcenter.blogspot.fr/2012/04/byte-order-fallacy.html
//
// This package could be useful ONLY to workaround bugs in other software.
//
//
// Note: reading godoc on one particular platform may be misleading.
// Check this:
//   GOARCH=amd64   godoc github.com/dolmen-go/endian
//   GOARCH=mips64  godoc github.com/dolmen-go/endian
//   GOARCH=unknown godoc github.com/dolmen-go/endian
package endian
