//ff:func feature=sql type=test control=sequence
//ff:what TestReceiverTypeName_UnknownCov 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestReceiverTypeName_UnknownCov(t *testing.T) {
	if receiverTypeName(&ast.CompositeLit{}) != "" {
		t.Fatal("expected empty")
	}
}
