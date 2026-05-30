//ff:func feature=scan type=test control=sequence topic=express
//ff:what recursiveParts 테스트 헬퍼
package express

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

// recursiveParts returns (innerCall, outerCall) for the .put level of
// router.route(p).get(h).put(h2).
func recursiveParts(t *testing.T, fi *fileInfo) (inner, outer *sitter.Node) {
	t.Helper()
	outer = outermostCall(fi)
	mem := findChildByType(outer, "member_expression")
	if mem == nil {
		t.Fatal("no member")
	}
	inner = mem.ChildByFieldName("object")
	return inner, outer
}
