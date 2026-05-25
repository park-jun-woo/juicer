//ff:func feature=sql type=parse control=sequence
//ff:what TestExtractParams_SkipContext 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestExtractParams_SkipContext(t *testing.T) {
	fields := &ast.FieldList{
		List: []*ast.Field{
			{Names: []*ast.Ident{{Name: "ctx"}}, Type: &ast.SelectorExpr{X: &ast.Ident{Name: "context"}, Sel: &ast.Ident{Name: "Context"}}},
		},
	}
	result := extractParams(fields)
	if len(result) != 0 {
		t.Fatalf("expected 0 (context skipped), got %d", len(result))
	}
}
