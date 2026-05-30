//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestUnwrapChainRecursive_NotHTTP 테스트
package express

import "testing"

func TestUnwrapChainRecursive_NotHTTP(t *testing.T) {

	fi := mustParse(t, []byte(`router.route('/x').get(g).foo(p);`))
	inner, outer := recursiveParts(t, fi)
	path, rv, methods := unwrapChainRecursive(inner, outer, "foo", fi.Src, map[string]bool{"router": true})
	if path != "/x" || rv != "router" || len(methods) != 1 {
		t.Fatalf("expected inner only, got path=%q methods=%v", path, methods)
	}
}
