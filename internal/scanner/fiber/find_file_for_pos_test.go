//ff:func feature=scan type=test control=sequence
//ff:what findFileForPos — 위치 기반 파일 검색 테스트
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestFindFileForPos(t *testing.T) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", "package m\nfunc F() {}\n", 0)
	if err != nil {
		t.Fatal(err)
	}
	pkg := &packages.Package{Syntax: []*ast.File{file}}
	pkgs := []*packages.Package{pkg}

	// a position inside the file resolves to it
	if got := findFileForPos(file.Pos()+1, pkgs); got != file {
		t.Fatalf("expected to find file for in-range pos")
	}

	// a position outside any file -> nil
	if got := findFileForPos(file.End()+100, pkgs); got != nil {
		t.Fatalf("expected nil for out-of-range pos, got %v", got)
	}
}
