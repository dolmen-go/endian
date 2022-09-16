//go:build !endiangen && go1.17 && !go1.18
// +build !endiangen,go1.17,!go1.18

//go:generate go run -tags endiangen generate_go1.17.go

package endian
