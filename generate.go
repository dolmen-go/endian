// To run this program: go generate
//
// (the go:generate line is in doc.go)
//
// +build generate

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/dolmen-go/codegen"
)

// parseArch extracts the "BigEndian" const from an arch definition source
// $GOROOT/src/runtime/internal/sys
func parseArch(filename string) (bool, error) {
	fs := token.NewFileSet()
	fileAST, err := parser.ParseFile(fs, filename, nil, parser.Mode(0))
	//fileAST, err := parser.ParseFile(fs, filename, nil, parser.Trace)
	if err != nil {
		return false, err
	}

	if len(fileAST.Decls) == 0 {
		return false, fmt.Errorf("%s: no Decls in AST", filename)
	}
	decl, ok := fileAST.Decls[0].(*ast.GenDecl)
	if !ok || decl.Tok != token.CONST {
		return false, fmt.Errorf("%s: CONST expected at Decls[0]", filename)
	}
	for _, rawSpec := range decl.Specs {
		spec := rawSpec.(*ast.ValueSpec)
		if len(spec.Names) != 1 || spec.Names[0].Name != "BigEndian" {
			continue
		}

		// We found the const "BigEndian"
		// Let's extract its value!
		if len(spec.Values) != 1 {
			return false, fmt.Errorf("%s: single value expected for const BigEndian", filename)
		}
		valueExpr, ok := spec.Values[0].(*ast.BasicLit)
		// fmt.Printf("%#v\n", valueExpr)
		if !ok {
			return false, fmt.Errorf("%s: BasicLit value expected for const BigEndian")
		}
		if valueExpr.Kind != token.INT {
			return false, fmt.Errorf("%s: INT value expected for const BigEndian")
		}

		intValue, _ := strconv.ParseInt(valueExpr.Value, 0, 64)
		if intValue < 0 || intValue > 1 {
			return false, fmt.Errorf("%s: value 0/1 expected for const BigEndian")
		}
		return intValue == 1, nil
	}

	return true, fmt.Errorf("%s: const BigEndian not found", filename)
}

func main() {
	var srcDir = os.ExpandEnv("${GOROOT}/src/runtime/internal/sys")
	dir, err := os.Open(srcDir)
	if err != nil {
		log.Fatal(err)
	}
	defer dir.Close()

	files, err := dir.Readdirnames(0)
	if err != nil {
		log.Fatal(err)
	}

	var knownArchs []string
	archsByEndian := make(map[bool][]string)

	for _, f := range files {
		if !strings.HasPrefix(f, "arch_") || !strings.HasSuffix(f, ".go") {
			continue
		}
		arch := f[5 : len(f)-3]
		//fmt.Println(arch)
		big, err := parseArch(srcDir + "/" + f)
		if err != nil {
			log.Println(err)
			continue
		}
		archsByEndian[big] = append(archsByEndian[big], arch)
		knownArchs = append(knownArchs, arch)
	}

	//fmt.Printf("%#v\n", archByEndian)

	const template = `
{{- /**/}}// +build {{- range .Tags}} {{.}}{{end}}

package endian

import "encoding/binary"

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

	// FIXME Remove ",!others"
	const templateOthers = `
{{- /**/}}// +build {{range .}}!{{.}},{{end}}!others

package endian

import (
	"encoding/binary"
	"unsafe"
)

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