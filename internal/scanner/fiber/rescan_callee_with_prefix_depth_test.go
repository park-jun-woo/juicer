//ff:func feature=scan type=test control=sequence
//ff:what rescanCalleeWithPrefixDepth — depth/해석 실패 분기 테스트
package fiber

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestRescanCalleeWithPrefixDepth_MaxDepth(t *testing.T) {
	call := parseCall(t, "registerRoutes(app)")
	ctx := &groupArgCtx{idx: &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}}, info: newEmptyInfo()}
	// depth at the limit -> immediate return, no panic
	rescanCalleeWithPrefixDepth(call, 0, "/api", &routerInfo{}, ctx, maxRescanDepth)
}

func TestRescanCalleeWithPrefixDepth_InvalidTarget(t *testing.T) {
	call := parseCall(t, "registerRoutes(app)")
	ctx := &groupArgCtx{idx: &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}}, info: newEmptyInfo()}
	// empty info -> resolveCallTarget invalid -> return
	rescanCalleeWithPrefixDepth(call, 0, "/api", &routerInfo{}, ctx, 0)
}

func TestRescanCalleeWithPrefixDepth_ResolvedNonRouterParam(t *testing.T) {
	// type-check a package where a call resolves to a local func, but the
	// callee's param is not a fiber router -> reaches lookupFunc + param check,
	// returns at the paramName == "" guard.
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

	// build index: map the registerRoutes func decl by its name's position
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
	if call == nil {
		t.Fatal("call not found")
	}
	ctx := &groupArgCtx{idx: idx, info: info, fset: fset}
	// resolves to registerRoutes (body present), but param "x int" is not a
	// router -> fiberRouterParamAtIndex "" -> return. No panic.
	rescanCalleeWithPrefixDepth(call, 0, "/api", &routerInfo{}, ctx, 0)
}
