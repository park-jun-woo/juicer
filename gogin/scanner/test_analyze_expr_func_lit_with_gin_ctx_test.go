//ff:func feature=scan type=extract control=sequence
//ff:what TestAnalyzeExpr_FuncLit_WithGinCtx 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestAnalyzeExpr_FuncLit_WithGinCtx(t *testing.T) {
	// FuncLit with *gin.Context param — should call scanBody
	funcLit := &ast.FuncLit{
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{{Name: "c"}},
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
	ep := &Endpoint{}
	info := &types.Info{
		Uses:       make(map[*ast.Ident]types.Object),
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
	}
	idx := &funcIndex{
		byPos: make(map[token.Pos]*ast.FuncDecl),
		info:  make(map[token.Pos]*types.Info),
	}
	analyzeExpr(ep, funcLit, info, idx)
	// No crash expected, scanBody called with empty body
}
