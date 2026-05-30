//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractRouteCallPath_NoString 테스트
package express

import "testing"

func TestExtractRouteCallPath_NoString(t *testing.T) {
	fi := mustParse(t, []byte(`router.route(p);`))
	if got := extractRouteCallPath(firstCallExpr(t, fi), fi.Src); got != "" {
		t.Fatalf("got %q", got)
	}
}
