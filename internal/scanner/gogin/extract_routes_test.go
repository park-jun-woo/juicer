//ff:func feature=scan type=test control=sequence
//ff:what TestExtractRoutes_NilPkgs 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestExtractRoutes_NilPkgs(t *testing.T) {
	result, _ := extractRoutes(nil, ".")
	if len(result) != 0 {
		t.Fatalf("expected 0, got %d", len(result))
	}

	// pkg with syntax but empty files
	fset := token.NewFileSet()
	f := fset.AddFile("test.go", -1, 100)
	_ = f
	file := &ast.File{Name: &ast.Ident{Name: "main"}}
	pkg := &packages.Package{
		Syntax:          []*ast.File{file},
		CompiledGoFiles: []string{"/tmp/test.go"},
		Fset:            fset,
	}
	result2, _ := extractRoutes([]*packages.Package{pkg}, "/tmp")
	_ = result2
}

