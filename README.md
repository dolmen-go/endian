# endian.Native - A single constant that you should not use

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/dolmen-go/endian)
[![Codecov](https://codecov.io/gh/dolmen-go/endian/graph/badge.svg?token=AUuPGQ01UE)](https://codecov.io/gh/dolmen-go/endian)[![Go Report Card](https://goreportcard.com/badge/github.com/dolmen-go/endian)](https://goreportcard.com/report/github.com/dolmen-go/endian)

## WARNING 1

You should probably not use this package, because you should avoid using the native byte ordering in programs!

Read this article by Rob Pike before using this package:

https://commandcenter.blogspot.fr/2012/04/byte-order-fallacy.html

You should probably first try to fix the broken code that generates
data dependent on the architecture on which it is compiled. Fix it so it
always generates code using a fixed byte order.
If you can't, see if you can detect the byte order of the data from the
data itself.


## WARNING 2

Go 1.21 has added [NativeEndian](https://pkg.go.dev/encoding/binary#NativeEndian) (see proposal [#57237](https://go.dev/issue/57237)) which serves the same purpose. However that implementation has a bug: you can't use `==` to compare `binary.NativeEndian` with either `binary.BigEndian` or `binary.LittleEndian`. See [#67026](https://go.dev/issue/67026) and the proposed fix in [CL 581655](https://go.dev/cl/581655).

This package doesn't have this bug, and is compatible even with very old versions of Go.

## Usage

This package only exports a single variable containing the [byte order](https://pkg.go.dev/encoding/binary#ByteOrder) of
GOARCH.

See the [encoding/binary](https://pkg.go.dev/encoding/binary) package
for how to use it to read/write a binary data stream.

```go
package endian

import "encoding/binary"

var Native binary.ByteOrder
```

The implementation is ultra lightweight because it relies on build tags.

## License

Copyright 2016-2022 Olivier Mengué

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
