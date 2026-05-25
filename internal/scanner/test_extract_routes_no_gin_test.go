//ff:func feature=scan type=extract control=sequence
//ff:what TestExtractRoutes_NoGin 테스트
package scanner

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestExtractRoutes_NoGin(t *testing.T) {
	src := `package main
func main() {}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "main.go", src, 0)

	pkg := &packages.Package{
		Syntax:          []*ast.File{file},
		CompiledGoFiles: []string{"main.go"},
		Fset:            fset,
	}

	eps := extractRoutes([]*packages.Package{pkg}, ".")
	if len(eps) != 0 {
		t.Errorf("expected 0 endpoints for non-gin code, got %d", len(eps))
	}
}
