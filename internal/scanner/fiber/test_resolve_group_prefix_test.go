//ff:func feature=scan type=test control=sequence
//ff:what TestResolveGroupPrefix 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"testing"
)

func TestResolveGroupPrefix(t *testing.T) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "/proj/m.go", "package m\nfunc F() {}\n", 0)
	if err != nil {
		t.Fatal(err)
	}
	pkgs := []*packages.Package{

		{Syntax: []*ast.File{file}, TypesInfo: nil, CompiledGoFiles: []string{"/proj/m.go"}, Fset: fset},

		{Syntax: []*ast.File{file}, TypesInfo: &types.Info{}, CompiledGoFiles: nil, Fset: fset},
	}
	idx := &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}}

	resolveGroupPrefix(pkgs, "/proj", idx, []scanner.Endpoint{}, map[int][]ast.Expr{})
}
