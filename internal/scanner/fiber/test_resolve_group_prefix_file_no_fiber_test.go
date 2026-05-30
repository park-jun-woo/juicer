//ff:func feature=scan type=test control=sequence
//ff:what TestResolveGroupPrefixFile_NoFiber 테스트
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

func TestResolveGroupPrefixFile_NoFiber(t *testing.T) {
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "/proj/m.go", "package m\nfunc F() {}\n", 0)
	pkg := &packages.Package{Syntax: []*ast.File{file}, TypesInfo: &types.Info{}, Fset: fset}
	idx := &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}}

	resolveGroupPrefixFile(file, pkg, []*packages.Package{pkg}, "/proj", idx, []scanner.Endpoint{}, map[int][]ast.Expr{}, epIndexEmpty())
}
