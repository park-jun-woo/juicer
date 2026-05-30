//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestUnwrapChainRecursive_AddsMethod 테스트
package express

import "testing"

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
