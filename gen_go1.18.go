//go:build !endiangen && go1.18 && !go1.19
// +build !endiangen,go1.18,!go1.19

//go:generate go run -tags endiangen generate_go1.18.go

package endian
