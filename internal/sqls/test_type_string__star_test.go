//ff:func feature=sql type=parse control=sequence
//ff:what TestTypeString_Star 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestTypeString_Star(t *testing.T) {
	got := typeString(&ast.StarExpr{X: &ast.Ident{Name: "int"}})
	if got != "*int" {
		t.Fatalf("expected *int, got %s", got)
	}
}
