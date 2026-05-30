//ff:func feature=scan type=test control=sequence topic=express
//ff:what unwrapChainBase: 경로+메서드 추출 / 경로없음 / 비http메서드 분기
package express

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
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

func TestUnwrapChainBase_Valid(t *testing.T) {
	fi := mustParse(t, []byte(`router.route('/x').get(h);`))
	route, outer := chainParts(t, fi)
	path, methods := unwrapChainBase(route, outer, "get", fi.Src)
	if path != "/x" || len(methods) != 1 || methods[0].method != "GET" {
		t.Fatalf("path=%q methods=%v", path, methods)
	}
}

func TestUnwrapChainBase_NoPath(t *testing.T) {
	fi := mustParse(t, []byte(`router.route(varPath).get(h);`))
	route, outer := chainParts(t, fi)
	if path, m := unwrapChainBase(route, outer, "get", fi.Src); path != "" || m != nil {
		t.Fatalf("expected empty, got %q %v", path, m)
	}
}

func TestUnwrapChainBase_NotHTTPMethod(t *testing.T) {
	fi := mustParse(t, []byte(`router.route('/x').foo(h);`))
	route, outer := chainParts(t, fi)
	if path, m := unwrapChainBase(route, outer, "foo", fi.Src); path != "" || m != nil {
		t.Fatalf("expected empty, got %q %v", path, m)
	}
}
