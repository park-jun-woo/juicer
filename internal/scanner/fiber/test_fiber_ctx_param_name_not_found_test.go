//ff:func feature=scan type=test control=sequence
//ff:what TestFiberCtxParamName_NotFound 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestFiberCtxParamName_NotFound(t *testing.T) {
	ft := &ast.FuncType{
		Params: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{{Name: "w"}},
					Type:  &ast.Ident{Name: "ResponseWriter"},
				},
			},
		},
	}
	got := fiberCtxParamName(ft)
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
