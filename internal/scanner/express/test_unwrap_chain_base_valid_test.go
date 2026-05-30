//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestUnwrapChainBase_Valid 테스트
package express

import "testing"

func TestUnwrapChainBase_Valid(t *testing.T) {
	fi := mustParse(t, []byte(`router.route('/x').get(h);`))
	route, outer := chainParts(t, fi)
	path, methods := unwrapChainBase(route, outer, "get", fi.Src)
	if path != "/x" || len(methods) != 1 || methods[0].method != "GET" {
		t.Fatalf("path=%q methods=%v", path, methods)
	}
}
