//go:build !generate && go1.17 && !go1.18
// +build !generate,go1.17,!go1.18

//go:generate go run -tags generate generate_go1.17.go

package endian
