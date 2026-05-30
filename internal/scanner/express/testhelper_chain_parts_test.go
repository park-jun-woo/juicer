//ff:func feature=scan type=test control=sequence topic=express
//ff:what chainParts 테스트 헬퍼
package express

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

// chainParts returns (routeCall, outerCall) for `router.route(path).get(h)`.
func chainParts(t *testing.T, fi *fileInfo) (route, outer *sitter.Node) {
	t.Helper()
	outer = outermostCall(fi)
	mem := findChildByType(outer, "member_expression")
	if mem == nil {
		t.Fatal("no member")
	}
	route = mem.ChildByFieldName("object")
	return route, outer
}
