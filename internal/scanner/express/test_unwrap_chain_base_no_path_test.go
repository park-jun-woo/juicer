//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestUnwrapChainBase_NoPath 테스트
package express

import "testing"

func TestUnwrapChainBase_NoPath(t *testing.T) {
	fi := mustParse(t, []byte(`router.route(varPath).get(h);`))
	route, outer := chainParts(t, fi)
	if path, m := unwrapChainBase(route, outer, "get", fi.Src); path != "" || m != nil {
		t.Fatalf("expected empty, got %q %v", path, m)
	}
}
