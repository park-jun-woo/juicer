//ff:func feature=scan type=extract control=sequence
//ff:what TestExtractRoutes_MoreSyntaxThanCompiled 테스트
package gogin

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestExtractRoutes_MoreSyntaxThanCompiled(t *testing.T) {
	src := `package main
func main() {}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "main.go", src, 0)

	// More Syntax files than CompiledGoFiles — should skip extras
	pkg := &packages.Package{
		Syntax:          []*ast.File{file, file},
		CompiledGoFiles: []string{"main.go"},
		Fset:            fset,
	}

	eps, _ := extractRoutes([]*packages.Package{pkg}, ".")
	if len(eps) != 0 {
		t.Errorf("expected 0, got %d", len(eps))
	}
}
