//ff:func feature=scan type=test control=sequence
//ff:what TestScanBody_WithJSONCallCov 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestScanBody_WithJSONCallCov(t *testing.T) {
	ep := &Endpoint{}
	idx := &funcIndex{byPos: make(map[token.Pos]*ast.FuncDecl), info: make(map[token.Pos]*types.Info)}
	call := &ast.CallExpr{
		Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "c"}, Sel: &ast.Ident{Name: "JSON"}},
		Args: []ast.Expr{
			&ast.BasicLit{Kind: token.INT, Value: "200"},
			&ast.Ident{Name: "data"},
		},
	}
	body := &ast.BlockStmt{List: []ast.Stmt{&ast.ExprStmt{X: call}}}
	info := &types.Info{Uses: make(map[*ast.Ident]types.Object), Types: make(map[ast.Expr]types.TypeAndValue), Selections: make(map[*ast.SelectorExpr]*types.Selection)}
	scanBody(ep, body, "c", info, idx, "handler")
	if len(ep.Responses) == 0 {
		t.Fatal("expected response from c.JSON")
	}
}
