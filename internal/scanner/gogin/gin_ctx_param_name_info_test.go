//ff:func feature=scan type=test control=sequence
//ff:what TestGinCtxParamNameInfo_NilInfo 테스트
package gogin

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestGinCtxParamNameInfo_NilInfo(t *testing.T) {
	ginCtxField := &ast.Field{
		Names: []*ast.Ident{{Name: "c"}},
		Type: &ast.StarExpr{
			X: &ast.SelectorExpr{
				X:   &ast.Ident{Name: "gin"},
				Sel: &ast.Ident{Name: "Context"},
			},
		},
	}
	ft := &ast.FuncType{
		Params: &ast.FieldList{List: []*ast.Field{ginCtxField}},
	}

	// nil info -> AST fallback
	got := ginCtxParamNameInfo(ft, nil)
	if got != "c" {
		t.Fatalf("expected c (AST fallback), got %s", got)
	}

	// info with non-gin type -> fallback to AST
	info := &types.Info{
		Types: map[ast.Expr]types.TypeAndValue{
			ginCtxField.Type: {Type: types.Typ[types.Int]},
		},
	}
	got = ginCtxParamNameInfo(ft, info)
	if got != "c" {
		t.Fatalf("expected c (AST fallback), got %s", got)
	}

	// info with no type info -> TypeOf returns nil, fallback
	info2 := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}
	got = ginCtxParamNameInfo(ft, info2)
	if got != "c" {
		t.Fatalf("expected c (AST fallback nil type), got %s", got)
	}
}


func TestGinCtxParamNameInfo_InfoFallback(t *testing.T) {
	// non-nil info but TypeOf returns nil for synthetic node -> AST fallback
	ft := &ast.FuncType{
		Params: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{{Name: "c"}},
					Type: &ast.StarExpr{X: &ast.SelectorExpr{
						X:   &ast.Ident{Name: "gin"},
						Sel: &ast.Ident{Name: "Context"},
					}},
				},
			},
		},
	}
	info := &types.Info{}
	if got := ginCtxParamNameInfo(ft, info); got != "c" {
		t.Fatalf("fallback: got %q, want c", got)
	}
}

func TestGinCtxParamNameInfo_NonGinFallback(t *testing.T) {
	ft := &ast.FuncType{
		Params: &ast.FieldList{
			List: []*ast.Field{
				{Names: []*ast.Ident{{Name: "x"}}, Type: &ast.Ident{Name: "int"}},
			},
		},
	}
	if got := ginCtxParamNameInfo(ft, &types.Info{}); got != "" {
		t.Fatalf("non-gin: got %q", got)
	}
}
