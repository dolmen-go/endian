//go:build !generate
// +build !generate

// Package endian exposes the native byte order of GOARCH.
//
// You should not use this package before reading and understanding this
// article by Rob Pike:
// https://commandcenter.blogspot.fr/2012/04/byte-order-fallacy.html
//
// This package could be useful ONLY to workaround bugs in other software.
//
// Note: reading godoc on one particular platform may be misleading.
// Check this:
//
//	GOARCH=amd64   go doc github.com/dolmen-go/endian Native
//	GOARCH=mips64  go doc github.com/dolmen-go/endian Native
//	GOARCH=unknown go doc github.com/dolmen-go/endian Native
package endian
