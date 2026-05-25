//ff:func feature=sql type=parse control=sequence
//ff:what TestReceiverTypeName_Star 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestReceiverTypeName_Star(t *testing.T) {
	expr := &ast.StarExpr{X: &ast.Ident{Name: "UserRepo"}}
	got := receiverTypeName(expr)
	if got != "UserRepo" {
		t.Fatalf("expected UserRepo, got %s", got)
	}
}
