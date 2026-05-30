//ff:func feature=scan type=test control=sequence
//ff:what TestAnalyzeExpr_FuncLit 테스트
package gogin

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"go/token"
	"go/types"
	"testing"
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

	sel := &ast.SelectorExpr{X: &ast.Ident{Name: "h"}, Sel: &ast.Ident{Name: "Method"}}
	info2 := &types.Info{
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
		Uses:       make(map[*ast.Ident]types.Object),
	}
	analyzeExpr(ep, sel, info2, idx)

	ident := &ast.Ident{Name: "handler"}
	info3 := &types.Info{
		Uses: make(map[*ast.Ident]types.Object),
	}
	analyzeExpr(ep, ident, info3, idx)

	call := &ast.CallExpr{
		Fun: &ast.Ident{Name: "getHandler"},
	}
	analyzeExpr(ep, call, info3, idx)

	analyzeExpr(ep, &ast.BasicLit{}, info, idx)

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
