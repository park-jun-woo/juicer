//ff:func feature=scan type=test control=sequence topic=express
//ff:what extractRouteCallPath: 정상 / args없음 / string없음 분기
package express

import "testing"

func TestExtractRouteCallPath_Valid(t *testing.T) {
	fi := mustParse(t, []byte(`router.route('/users');`))
	if got := extractRouteCallPath(firstCallExpr(t, fi), fi.Src); got != "/users" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractRouteCallPath_NoArgs(t *testing.T) {
	fi := mustParse(t, []byte("router.route`x`;"))
	if got := extractRouteCallPath(firstCallExpr(t, fi), fi.Src); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractRouteCallPath_NoString(t *testing.T) {
	fi := mustParse(t, []byte(`router.route(p);`))
	if got := extractRouteCallPath(firstCallExpr(t, fi), fi.Src); got != "" {
		t.Fatalf("got %q", got)
	}
}
