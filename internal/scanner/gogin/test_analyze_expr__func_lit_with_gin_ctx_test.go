//ff:func feature=scan type=extract control=sequence
//ff:what TestAnalyzeExpr_FuncLitWithGinCtx 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestAnalyzeExpr_FuncLitWithGinCtx(t *testing.T) {
	ep := &scanner.Endpoint{}
	fn := &ast.FuncLit{
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
	info := &types.Info{
		Uses:       make(map[*ast.Ident]types.Object),
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
	}
	idx := &funcIndex{byPos: make(map[token.Pos]*ast.FuncDecl), info: make(map[token.Pos]*types.Info)}
	analyzeExpr(ep, fn, info, idx)
}
