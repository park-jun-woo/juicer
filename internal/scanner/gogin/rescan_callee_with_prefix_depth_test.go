//ff:func feature=scan type=test control=sequence
//ff:what rescanCalleeWithPrefixDepth — depth/해석 분기 테스트
package gogin

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func goginEmptyInfo() *types.Info {
	return &types.Info{
		Uses:       map[*ast.Ident]types.Object{},
		Selections: map[*ast.SelectorExpr]*types.Selection{},
	}
}

func goginParseCall(t *testing.T, expr string) *ast.CallExpr {
	t.Helper()
	e, err := parser.ParseExpr(expr)
	if err != nil {
		t.Fatal(err)
	}
	return e.(*ast.CallExpr)
}

func TestRescanCalleeWithPrefixDepth_MaxDepth(t *testing.T) {
	call := goginParseCall(t, "registerRoutes(r)")
	ctx := &groupArgCtx{idx: &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}}, info: goginEmptyInfo()}
	rescanCalleeWithPrefixDepth(call, 0, "/api", &routerInfo{}, ctx, maxRescanDepth)
}

func TestRescanCalleeWithPrefixDepth_InvalidTarget(t *testing.T) {
	call := goginParseCall(t, "registerRoutes(r)")
	ctx := &groupArgCtx{idx: &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}}, info: goginEmptyInfo()}
	rescanCalleeWithPrefixDepth(call, 0, "/api", &routerInfo{}, ctx, 0)
}

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
	idx := &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}, info: map[token.Pos]*types.Info{}}
	var call *ast.CallExpr
	for _, d := range file.Decls {
		fn, ok := d.(*ast.FuncDecl)
		if !ok {
			continue
		}
		idx.byPos[fn.Name.Pos()] = fn
		idx.info[fn.Name.Pos()] = info
		ast.Inspect(fn, func(n ast.Node) bool {
			if c, ok := n.(*ast.CallExpr); ok && call == nil {
				if id, ok := c.Fun.(*ast.Ident); ok && id.Name == "registerRoutes" {
					call = c
				}
			}
			return true
		})
	}
	ctx := &groupArgCtx{idx: idx, info: info, fset: fset}
	rescanCalleeWithPrefixDepth(call, 0, "/api", &routerInfo{}, ctx, 0)
}

func TestRescanCalleeWithPrefixDepth_NilParent(t *testing.T) {
	// depth limit hit immediately with a nil parent -> no deref, no panic
	call := goginParseCall(t, "f(r)")
	ctx := &groupArgCtx{idx: &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}}, info: goginEmptyInfo()}
	rescanCalleeWithPrefixDepth(call, 0, "/x", &routerInfo{}, ctx, maxRescanDepth+1)
}
