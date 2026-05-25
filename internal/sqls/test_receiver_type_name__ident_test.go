//ff:func feature=sql type=parse control=sequence
//ff:what TestReceiverTypeName_Ident 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestReceiverTypeName_Ident(t *testing.T) {
	got := receiverTypeName(&ast.Ident{Name: "UserRepo"})
	if got != "UserRepo" {
		t.Fatalf("expected UserRepo, got %s", got)
	}
}
