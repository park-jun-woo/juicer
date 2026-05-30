//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractRouteCallPath_Valid 테스트
package express

import "testing"

func TestExtractRouteCallPath_Valid(t *testing.T) {
	fi := mustParse(t, []byte(`router.route('/users');`))
	if got := extractRouteCallPath(firstCallExpr(t, fi), fi.Src); got != "/users" {
		t.Fatalf("got %q", got)
	}
}
