//ff:func feature=sql type=parse control=sequence
//ff:what TestReceiverTypeName_Unknown 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestReceiverTypeName_Unknown(t *testing.T) {
	got := receiverTypeName(&ast.BasicLit{})
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
