//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractPathRouteFromObject_Valid 테스트
package express

import "testing"

func TestExtractPathRouteFromObject_Valid(t *testing.T) {
	fi := mustParse(t, []byte(`const o = { path: '/a', route: r };`))
	e := extractPathRouteFromObject(firstObject(t, fi), fi.Src)
	if e == nil || e.path != "/a" || e.routeVar != "r" {
		t.Fatalf("got %+v", e)
	}
}
