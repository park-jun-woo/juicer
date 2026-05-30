//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractRouteCallPath_NoArgs 테스트
package express

import "testing"

func TestExtractRouteCallPath_NoArgs(t *testing.T) {
	fi := mustParse(t, []byte("router.route`x`;"))
	if got := extractRouteCallPath(firstCallExpr(t, fi), fi.Src); got != "" {
		t.Fatalf("got %q", got)
	}
}
