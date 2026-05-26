//ff:func feature=sql type=test control=sequence
//ff:what TestExtractParams_WithNamesCov 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestExtractParams_WithNamesCov(t *testing.T) {
	fields := &ast.FieldList{List: []*ast.Field{
		{Names: []*ast.Ident{{Name: "ctx"}}, Type: &ast.SelectorExpr{X: &ast.Ident{Name: "context"}, Sel: &ast.Ident{Name: "Context"}}},
		{Names: []*ast.Ident{{Name: "id"}}, Type: &ast.Ident{Name: "int"}},
		{Type: &ast.Ident{Name: "string"}},
	}}
	params := extractParams(fields)
	if len(params) < 1 {
		t.Fatal("expected at least 1 param")
	}
}
