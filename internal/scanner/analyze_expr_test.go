package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestAnalyzeExpr_FuncLit(t *testing.T) {
	ep := &Endpoint{}
	fn := &ast.FuncLit{
		Type: &ast.FuncType{Params: &ast.FieldList{}},
		Body: &ast.BlockStmt{},
	}
	info := &types.Info{Uses: make(map[*ast.Ident]types.Object)}
	idx := &funcIndex{byPos: make(map[token.Pos]*ast.FuncDecl), info: make(map[token.Pos]*types.Info)}
	analyzeExpr(ep, fn, info, idx)
}

func TestAnalyzeExpr_FuncLitWithGinCtx(t *testing.T) {
	ep := &Endpoint{}
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

func TestAnalyzeExpr_SelectorExprNoSelection(t *testing.T) {
	ep := &Endpoint{}
	sel := &ast.SelectorExpr{
		X:   &ast.Ident{Name: "h"},
		Sel: &ast.Ident{Name: "Method"},
	}
	info := &types.Info{
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
		Uses:       make(map[*ast.Ident]types.Object),
	}
	idx := &funcIndex{byPos: make(map[token.Pos]*ast.FuncDecl), info: make(map[token.Pos]*types.Info)}
	analyzeExpr(ep, sel, info, idx)
}

func TestAnalyzeExpr_Ident(t *testing.T) {
	ep := &Endpoint{}
	ident := &ast.Ident{Name: "handler"}
	info := &types.Info{Uses: make(map[*ast.Ident]types.Object)}
	idx := &funcIndex{byPos: make(map[token.Pos]*ast.FuncDecl), info: make(map[token.Pos]*types.Info)}
	analyzeExpr(ep, ident, info, idx)
}

func TestAnalyzeExpr_CallExprCase(t *testing.T) {
	ep := &Endpoint{}
	call := &ast.CallExpr{
		Fun: &ast.Ident{Name: "handler"},
	}
	info := &types.Info{Uses: make(map[*ast.Ident]types.Object)}
	idx := &funcIndex{byPos: make(map[token.Pos]*ast.FuncDecl), info: make(map[token.Pos]*types.Info)}
	analyzeExpr(ep, call, info, idx)
}

func TestAnalyzeExpr_Nil(t *testing.T) {
	ep := &Endpoint{}
	info := &types.Info{Uses: make(map[*ast.Ident]types.Object)}
	idx := &funcIndex{byPos: make(map[token.Pos]*ast.FuncDecl), info: make(map[token.Pos]*types.Info)}
	analyzeExpr(ep, nil, info, idx)
}
