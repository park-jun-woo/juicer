//ff:func feature=sql type=parse control=sequence
//ff:what TestExtractParams 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestExtractParams(t *testing.T) {
	t.Run("nil fields", func(t *testing.T) {
		got := extractParams(nil)
		if got != nil {
			t.Error("expected nil")
		}
	})

	t.Run("with context", func(t *testing.T) {
		fields := &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{{Name: "ctx"}},
					Type: &ast.SelectorExpr{
						X:   &ast.Ident{Name: "context"},
						Sel: &ast.Ident{Name: "Context"},
					},
				},
				{
					Names: []*ast.Ident{{Name: "name"}},
					Type:  &ast.Ident{Name: "string"},
				},
			},
		}
		got := extractParams(fields)
		if len(got) != 1 {
			t.Errorf("expected 1 param (ctx skipped), got %d: %v", len(got), got)
		}
	})

	t.Run("unnamed param", func(t *testing.T) {
		fields := &ast.FieldList{
			List: []*ast.Field{
				{Type: &ast.Ident{Name: "int"}},
			},
		}
		got := extractParams(fields)
		if len(got) != 1 || got[0] != "int" {
			t.Errorf("expected ['int'], got %v", got)
		}
	})
}
