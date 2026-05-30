//ff:func feature=scan type=test control=sequence
//ff:what TestFiberCtxParamNameInfo_InfoNonFiberFallback 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestFiberCtxParamNameInfo_InfoNonFiberFallback(t *testing.T) {

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
