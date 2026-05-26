//ff:func feature=scan type=test control=sequence
//ff:what findFileForPos 전 분기 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestFindFileForPos(t *testing.T) {
	// empty pkgs
	got := findFileForPos(token.Pos(1), nil)
	if got != nil {
		t.Fatal("expected nil for empty pkgs")
	}

	// with pkg but no matching pos
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
