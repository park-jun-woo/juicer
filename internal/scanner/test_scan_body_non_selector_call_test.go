//ff:func feature=scan type=extract control=sequence
//ff:what TestScanBody_NonSelectorCall 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestScanBody_NonSelectorCall(t *testing.T) {
	// A call that is not a selector expr (e.g., a plain function call)
	body := &ast.BlockStmt{
		List: []ast.Stmt{
			&ast.ExprStmt{
				X: &ast.CallExpr{
					Fun: &ast.Ident{Name: "someFunc"},
				},
			},
		},
	}
	ep := &Endpoint{}
	info := &types.Info{
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Uses:       make(map[*ast.Ident]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
	}
	idx := &funcIndex{
		byPos: make(map[token.Pos]*ast.FuncDecl),
		info:  make(map[token.Pos]*types.Info),
	}
	scanBody(ep, body, "c", info, idx, "handler")
}
