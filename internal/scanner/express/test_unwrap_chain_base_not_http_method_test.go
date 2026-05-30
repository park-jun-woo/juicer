//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestUnwrapChainBase_NotHTTPMethod 테스트
package express

import "testing"

func TestUnwrapChainBase_NotHTTPMethod(t *testing.T) {
	fi := mustParse(t, []byte(`router.route('/x').foo(h);`))
	route, outer := chainParts(t, fi)
	if path, m := unwrapChainBase(route, outer, "foo", fi.Src); path != "" || m != nil {
		t.Fatalf("expected empty, got %q %v", path, m)
	}
}
