//ff:func feature=scan type=test control=sequence
//ff:what TestFindFileForPos_Found 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
	"testing"
)

func TestFindFileForPos_Found(t *testing.T) {
	fset := token.NewFileSet()
	file, err := parserParseFileG(fset, "package m\nfunc F() {}\n")
	if err != nil {
		t.Fatal(err)
	}
	pkg := &packages.Package{Syntax: []*ast.File{file}}
	if got := findFileForPos(file.Pos()+1, []*packages.Package{pkg}); got != file {
		t.Fatal("expected to find file for in-range pos")
	}
}
