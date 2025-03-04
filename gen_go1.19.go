//go:build !endiangen && go1.19 && !go1.24
// +build !endiangen,go1.19,!go1.24

//go:generate go run -tags endiangen generate_go1.19.go

package endian
