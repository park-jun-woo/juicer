//ff:func feature=sql type=test control=sequence
//ff:what TestReceiverTypeName_IdentCov 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestReceiverTypeName_IdentCov(t *testing.T) {
	expr := &ast.Ident{Name: "Repo"}
	if receiverTypeName(expr) != "Repo" {
		t.Fatal("expected Repo")
	}
}
