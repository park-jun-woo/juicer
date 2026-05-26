//ff:func feature=scan type=test control=sequence
//ff:what TestExprName_Ident 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestExprName_Ident(t *testing.T) {
	got := exprName(&ast.Ident{Name: "handler"})
	if got != "handler" {
		t.Fatalf("expected handler, got %s", got)
	}
}
