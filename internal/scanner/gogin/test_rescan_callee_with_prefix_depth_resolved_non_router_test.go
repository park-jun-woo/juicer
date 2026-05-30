//ff:func feature=scan type=test control=sequence
//ff:what TestRescanCalleeWithPrefixDepth_ResolvedNonRouter 테스트
package gogin

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestRescanCalleeWithPrefixDepth_ResolvedNonRouter(t *testing.T) {
	src := `package m
func registerRoutes(x int) {}
func Setup() { registerRoutes(0) }
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{Uses: map[*ast.Ident]types.Object{}, Selections: map[*ast.SelectorExpr]*types.Selection{}}
	if _, err := conf.Check("m", fset, []*ast.File{file}, info); err != nil {
		t.Fatal(err)
	}
	idx, call := indexDeclsAndFindCall(file, info)
	ctx := &groupArgCtx{idx: idx, info: info, fset: fset}
	rescanCalleeWithPrefixDepth(call, 0, "/api", &routerInfo{}, ctx, 0)
}
