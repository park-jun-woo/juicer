//ff:func feature=scan type=test control=sequence
//ff:what TestFindFileForPos 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
	"testing"
)

func TestFindFileForPos(t *testing.T) {

	got := findFileForPos(token.Pos(1), nil)
	if got != nil {
		t.Fatal("expected nil for empty pkgs")
	}

	fset := token.NewFileSet()
	f := fset.AddFile("test.go", -1, 100)
	_ = f
	pkg := &packages.Package{
		Syntax: []*ast.File{},
	}
	got = findFileForPos(token.Pos(50), []*packages.Package{pkg})
	if got != nil {
		t.Fatal("expected nil for no matching file")
	}
}
