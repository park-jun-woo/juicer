//ff:func feature=scan type=test control=sequence
//ff:what TestAnalyzeExpr_FuncLit 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestAnalyzeExpr_FuncLit(t *testing.T) {
	ep := &scanner.Endpoint{}
	fn := &ast.FuncLit{
		Type: &ast.FuncType{Params: &ast.FieldList{}},
		Body: &ast.BlockStmt{},
	}
	info := &types.Info{Uses: make(map[*ast.Ident]types.Object)}
	idx := &funcIndex{byPos: make(map[token.Pos]*ast.FuncDecl), info: make(map[token.Pos]*types.Info)}
	analyzeExpr(ep, fn, info, idx)

	// SelectorExpr - no selection
	sel := &ast.SelectorExpr{X: &ast.Ident{Name: "h"}, Sel: &ast.Ident{Name: "Method"}}
	info2 := &types.Info{
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
		Uses:       make(map[*ast.Ident]types.Object),
	}
	analyzeExpr(ep, sel, info2, idx)

	// Ident - no uses
	ident := &ast.Ident{Name: "handler"}
	info3 := &types.Info{
		Uses: make(map[*ast.Ident]types.Object),
	}
	analyzeExpr(ep, ident, info3, idx)

	// CallExpr
	call := &ast.CallExpr{
		Fun: &ast.Ident{Name: "getHandler"},
	}
	analyzeExpr(ep, call, info3, idx)

	// nil/unsupported expr
	analyzeExpr(ep, &ast.BasicLit{}, info, idx)

	// Ident with valid Uses entry — obj != nil but lookupFunc returns nil
	ident2 := &ast.Ident{Name: "myHandler"}
	pkg := types.NewPackage("example.com/test", "test")
	sig := types.NewSignatureType(nil, nil, nil, nil, nil, false)
	fnObj := types.NewFunc(token.NoPos, pkg, "myHandler", sig)
	info5 := &types.Info{
		Uses: map[*ast.Ident]types.Object{
			ident2: fnObj,
		},
	}
	analyzeExpr(ep, ident2, info5, idx)
}

func ginCtxFuncLit(name string) *ast.FuncLit {
	return &ast.FuncLit{
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{{Name: name}},
						Type: &ast.StarExpr{
							X: &ast.SelectorExpr{
								X:   &ast.Ident{Name: "gin"},
								Sel: &ast.Ident{Name: "Context"},
							},
						},
					},
				},
			},
		},
		Body: &ast.BlockStmt{},
	}
}

func TestAnalyzeExpr_FuncLitWithCtx(t *testing.T) {
	ep := &scanner.Endpoint{}
	idx := &funcIndex{byPos: make(map[token.Pos]*ast.FuncDecl), info: make(map[token.Pos]*types.Info)}
	// nil info -> AST fallback resolves ctx name "c"; scanBody runs (empty body).
	analyzeExpr(ep, ginCtxFuncLit("c"), nil, idx)
}

func TestAnalyzeExpr_DefaultBasicLit(t *testing.T) {
	ep := &scanner.Endpoint{}
	idx := &funcIndex{byPos: make(map[token.Pos]*ast.FuncDecl), info: make(map[token.Pos]*types.Info)}
	analyzeExpr(ep, &ast.BasicLit{Kind: token.INT, Value: "1"}, nil, idx)
	if ep.Request != nil {
		t.Fatal("default case should be a no-op")
	}
}
