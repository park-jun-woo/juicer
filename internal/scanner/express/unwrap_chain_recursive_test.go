//ff:func feature=scan type=test control=sequence topic=express
//ff:what unwrapChainRecursive: 메서드 추가 / innerPath없음 / 비http메서드
package express

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
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

func TestUnwrapChainRecursive_AddsMethod(t *testing.T) {
	fi := mustParse(t, []byte(`router.route('/x').get(g).put(p);`))
	inner, outer := recursiveParts(t, fi)
	path, rv, methods := unwrapChainRecursive(inner, outer, "put", fi.Src, map[string]bool{"router": true})
	if path != "/x" || rv != "router" || len(methods) != 2 {
		t.Fatalf("path=%q rv=%q methods=%v", path, rv, methods)
	}
	if methods[1].method != "PUT" {
		t.Fatalf("expected PUT appended, got %v", methods[1].method)
	}
}

func TestUnwrapChainRecursive_InnerNoPath(t *testing.T) {
	// inner is not a route chain -> innerPath empty
	fi := mustParse(t, []byte(`foo().put(p);`))
	inner, outer := recursiveParts(t, fi)
	if path, _, m := unwrapChainRecursive(inner, outer, "put", fi.Src, map[string]bool{"router": true}); path != "" || m != nil {
		t.Fatalf("expected empty, got %q %v", path, m)
	}
}

func TestUnwrapChainRecursive_NotHTTP(t *testing.T) {
	// outer prop is not an http method -> return inner only
	fi := mustParse(t, []byte(`router.route('/x').get(g).foo(p);`))
	inner, outer := recursiveParts(t, fi)
	path, rv, methods := unwrapChainRecursive(inner, outer, "foo", fi.Src, map[string]bool{"router": true})
	if path != "/x" || rv != "router" || len(methods) != 1 {
		t.Fatalf("expected inner only, got path=%q methods=%v", path, methods)
	}
}
