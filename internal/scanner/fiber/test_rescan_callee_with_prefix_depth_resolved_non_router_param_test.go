//ff:func feature=scan type=test control=sequence
//ff:what TestRescanCalleeWithPrefixDepth_ResolvedNonRouterParam 테스트
package fiber

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestRescanCalleeWithPrefixDepth_ResolvedNonRouterParam(t *testing.T) {

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
	info := &types.Info{
		Uses:       map[*ast.Ident]types.Object{},
		Selections: map[*ast.SelectorExpr]*types.Selection{},
	}
	if _, err := conf.Check("m", fset, []*ast.File{file}, info); err != nil {
		t.Fatal(err)
	}

	idx, call := indexDeclsAndFindCall(file, info)
	if call == nil {
		t.Fatal("call not found")
	}
	ctx := &groupArgCtx{idx: idx, info: info, fset: fset}

	rescanCalleeWithPrefixDepth(call, 0, "/api", &routerInfo{}, ctx, 0)
}
