// To run this program: go generate
//
// (the go:generate line is in gen_go1.19.go)
//
//go:build endiangen && go1.19
// +build endiangen,go1.19

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"runtime"
	"sort"
	"strings"

	"github.com/dolmen-go/codegen"
)

func readArchList() ([]string, error) {
	filename := runtime.GOROOT() + "/src/go/build/syslist.go"

	// This code is copied from go 1.19 $GOROOT/src/internal/goarch/gengoarch.go

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var goarches []string
	const (
		goarchVar    = `knownArch`
		goarchPrefix = `var ` + goarchVar + ` = map[string]bool{`
	)
	inGOARCH := false
	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, goarchPrefix) {
			inGOARCH = true
		} else if inGOARCH && strings.HasPrefix(line, "}") {
			break
		} else if inGOARCH {
			goarch := strings.Fields(line)[0]
			goarch = strings.TrimPrefix(goarch, `"`)
			goarch = strings.TrimSuffix(goarch, `":`)
			goarches = append(goarches, goarch)
		}
	}

	if len(goarches) == 0 {
		return nil, fmt.Errorf("%s: no "+goarchVar+" found", filename)
	}
	return goarches, nil
}

// getBigEndian extracts the list of "BigEndian" GOARCH values from $GOROOT/src/internal/goarch/goarch.go
func getBigEndian() ([]string, error) {
	filename := runtime.GOROOT() + "/src/internal/goarch/goarch.go"
	fs := token.NewFileSet()
	fileAST, err := parser.ParseFile(fs, filename, nil, parser.Mode(0))
	//fileAST, err := parser.ParseFile(fs, filename, nil, parser.Trace)
	if err != nil {
		return nil, err
	}

	/*

		// BigEndian reports whether the architecture is big-endian.
		const BigEndian = IsArmbe|IsArm64be|IsMips|IsMips64|IsPpc|IsPpc64|IsS390|IsS390x|IsSparc|IsSparc64 == 1

	*/

	if len(fileAST.Decls) == 0 {
		return nil, fmt.Errorf("%s: no Decls in AST", filename)
	}
	// fmt.Printf("%#v\n", fileAST.Decls)
	for _, decl := range fileAST.Decls {
		decl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		if decl.Tok != token.CONST {
			continue
		}
		spec := decl.Specs[0].(*ast.ValueSpec)
		if len(spec.Names) != 1 || spec.Names[0].Name != "BigEndian" {
			continue
		}
		// We found the const "BigEndian"
		// Let's extract its value!
		if len(spec.Values) != 1 {
			return nil, fmt.Errorf("%s: single value expected for const BigEndian", filename)
		}

		var archs []string

		list := spec.Values[0].(*ast.BinaryExpr).X.(*ast.BinaryExpr)
		for {
			arch := strings.ToLower(strings.TrimPrefix(list.Y.(*ast.Ident).Name, "Is"))
			archs = append(archs, arch)

			var ok bool
			list2, ok := list.X.(*ast.BinaryExpr)
			if !ok {
				arch = strings.ToLower(strings.TrimPrefix(list.X.(*ast.Ident).Name, "Is"))
				archs = append(archs, arch)
				break
			}
			list = list2
		}

		// Reverse
		for i, j := 0, len(archs)-1; i < j; i, j = i+1, j-1 {
			archs[i], archs[j] = archs[j], archs[i]
		}

		return archs, nil
	}

	return nil, fmt.Errorf("%s: const BigEndian not found", filename)
}

func main() {
	log.SetPrefix("")
	log.SetFlags(0)

	knownArchs, err := readArchList()
	if err != nil {
		log.Fatal(err)
	}

	bigEndianArchs, err := getBigEndian()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("knownArchs:", knownArchs)
	fmt.Println("bigEndianArchs:", bigEndianArchs)

	// Verify consistency
	for _, be := range bigEndianArchs {
		found := false
		for _, arch := range knownArchs {
			if be == arch {
				found = true
				break
			}
		}
		if !found {
			log.Fatalf("Goarch %s not found", be)
		}
	}

	archsByEndian := map[bool][]string{}

	for _, arch := range knownArchs {
		isBE := false
		for _, be := range bigEndianArchs {
			if be == arch {
				isBE = true
				break
			}
		}
		archsByEndian[isBE] = append(archsByEndian[isBE], arch)
	}

	// fmt.Printf("%#v\n", archsByEndian)

	const template = `// Code generated by generate_go1.19.go; DO NOT EDIT.

{{/**/}}//go:build ({{range $i, $tag := .Tags}}{{if gt $i 0}} || {{end}}{{$tag}}{{end}}) && !generate
{{/**/}}// +build {{- range .Tags}} {{.}}{{end}}
{{/**/}}// +build !generate

package endian

import "encoding/binary"

// Native is the byte order of GOARCH.
var Native = binary.{{if .Big}}Big{{else}}Little{{end}}Endian
`

	tmplArch := codegen.MustParse(template)

	bigStr := map[bool]string{true: "Big", false: "Little"}

	for big, tags := range archsByEndian {
		sort.Strings(tags)
		err = tmplArch.CreateFile(
			strings.ToLower(bigStr[big])+".go",
			map[string]interface{}{
				"Tags": tags,
				"Big":  big,
			},
		)
	}

	const templateOthers = `// Code generated by generate_go1.19.go; DO NOT EDIT.

{{/**/}}//go:build {{range .}}!{{.}} && {{end}}!generate
{{/**/}}// +build {{range .}}!{{.}},{{end}}!generate

package endian

import (
	"encoding/binary"
	"unsafe"
)

// Native is the byte order of GOARCH.
// It will be determined at runtime because it was unknown at code
// generation time.
var Native binary.ByteOrder

func init() {
	// http://grokbase.com/t/gg/golang-nuts/129jhmdb3d/go-nuts-how-to-tell-endian-ness-of-machine#20120918nttlyywfpl7ughnsys6pm4pgpe
	var x int32 = 0x01020304
	switch *(*byte)(unsafe.Pointer(&x)) {
	case 1:
		Native = binary.BigEndian
	case 4:
		Native = binary.LittleEndian
	}
}
`

	sort.Strings(knownArchs)

	if err = codegen.MustParse(templateOthers).CreateFile("others.go", knownArchs); err != nil {
		log.Fatal(err)
	}
}