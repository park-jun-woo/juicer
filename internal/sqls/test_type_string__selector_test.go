//ff:func feature=sql type=parse control=sequence
//ff:what TestTypeString_Selector 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestTypeString_Selector(t *testing.T) {
	got := typeString(&ast.SelectorExpr{X: &ast.Ident{Name: "sql"}, Sel: &ast.Ident{Name: "NullString"}})
	if got != "sql.NullString" {
		t.Fatalf("expected sql.NullString, got %s", got)
	}
}
