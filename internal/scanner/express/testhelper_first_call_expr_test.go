//ff:func feature=scan type=test control=sequence topic=express
//ff:what firstCallExpr 테스트 헬퍼
package express

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func firstCallExpr(t *testing.T, fi *fileInfo) *sitter.Node {
	t.Helper()
	calls := findAllByType(fi.Root, "call_expression")
	if len(calls) == 0 {
		t.Fatal("no call_expression")
	}
	return calls[0]
}
