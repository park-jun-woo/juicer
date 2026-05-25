//ff:func feature=sql type=parse control=sequence
//ff:what TestExtractReturns 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestExtractReturns(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		got := extractReturns(nil)
		if got != nil {
			t.Error("expected nil")
		}
	})

	t.Run("with named returns", func(t *testing.T) {
		fields := &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{{Name: "result"}},
					Type:  &ast.ArrayType{Elt: &ast.Ident{Name: "User"}},
				},
				{Type: &ast.Ident{Name: "error"}},
			},
		}
		got := extractReturns(fields)
		if len(got) != 2 {
			t.Errorf("expected 2 returns, got %d: %v", len(got), got)
		}
	})
}
