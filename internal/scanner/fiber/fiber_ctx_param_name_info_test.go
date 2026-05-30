//ff:func feature=scan type=test control=sequence
//ff:what fiberCtxParamNameInfo — types.Info 기반 ctx 파라미터 이름 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func ctxFuncType(name string) *ast.FuncType {
	return &ast.FuncType{
		Params: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{{Name: name}},
					Type: &ast.StarExpr{
						X: &ast.SelectorExpr{
							X:   &ast.Ident{Name: "fiber"},
							Sel: &ast.Ident{Name: "Ctx"},
						},
					},
				},
			},
		},
	}
}

func TestFiberCtxParamNameInfo_NilInfo(t *testing.T) {
	// info nil -> AST fallback path resolves "c"
	if got := fiberCtxParamNameInfo(ctxFuncType("c"), nil); got != "c" {
		t.Fatalf("nil info: got %q, want c", got)
	}
}

func TestFiberCtxParamNameInfo_NilParams(t *testing.T) {
	if got := fiberCtxParamNameInfo(&ast.FuncType{}, newEmptyInfo()); got != "" {
		t.Fatalf("nil params: got %q", got)
	}
}

func TestFiberCtxParamNameInfo_InfoNoMatchFallback(t *testing.T) {
	// non-nil info but TypeOf returns nil for our synthetic node ->
	// loop finds no fiber ctx via type info -> AST fallback recognizes it.
	if got := fiberCtxParamNameInfo(ctxFuncType("ctx"), newEmptyInfo()); got != "ctx" {
		t.Fatalf("fallback: got %q, want ctx", got)
	}
}

func TestFiberCtxParamNameInfo_InfoNonFiberFallback(t *testing.T) {
	// non-fiber param with non-nil info -> no type match, AST fallback yields ""
	ft := &ast.FuncType{
		Params: &ast.FieldList{
			List: []*ast.Field{
				{Names: []*ast.Ident{{Name: "x"}}, Type: &ast.Ident{Name: "int"}},
			},
		},
	}
	if got := fiberCtxParamNameInfo(ft, newEmptyInfo()); got != "" {
		t.Fatalf("expected empty for non-fiber param, got %q", got)
	}
}
