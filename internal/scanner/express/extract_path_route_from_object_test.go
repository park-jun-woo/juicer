//ff:func feature=scan type=test control=sequence topic=express
//ff:what extractPathRouteFromObject: 정상 / path누락 / route누락 분기
package express

import "testing"

func TestExtractPathRouteFromObject_Valid(t *testing.T) {
	fi := mustParse(t, []byte(`const o = { path: '/a', route: r };`))
	e := extractPathRouteFromObject(firstObject(t, fi), fi.Src)
	if e == nil || e.path != "/a" || e.routeVar != "r" {
		t.Fatalf("got %+v", e)
	}
}

func TestExtractPathRouteFromObject_MissingPath(t *testing.T) {
	fi := mustParse(t, []byte(`const o = { route: r };`))
	if e := extractPathRouteFromObject(firstObject(t, fi), fi.Src); e != nil {
		t.Fatalf("expected nil, got %+v", e)
	}
}

func TestExtractPathRouteFromObject_MissingRoute(t *testing.T) {
	fi := mustParse(t, []byte(`const o = { path: '/a' };`))
	if e := extractPathRouteFromObject(firstObject(t, fi), fi.Src); e != nil {
		t.Fatalf("expected nil, got %+v", e)
	}
}
