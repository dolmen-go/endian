// Code generated by generate_go1.18.go; DO NOT EDIT.

//go:build (386 || amd64 || amd64p32 || arm || arm64 || loong64 || mips64le || mips64p32 || mips64p32le || mipsle || ppc64le || riscv || riscv64 || wasm) && !generate
// +build 386 amd64 amd64p32 arm arm64 loong64 mips64le mips64p32 mips64p32le mipsle ppc64le riscv riscv64 wasm
// +build !generate

package endian

import "encoding/binary"

// Native is the byte order of GOARCH.
var Native = binary.LittleEndian
