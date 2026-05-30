//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractRouteFromCall_Nil 테스트
package express

import "testing"

func TestExtractRouteFromCall_Nil(t *testing.T) {
	fi := mustParse(t, []byte(`foo('/x');`))
	if routes := extractRouteFromCall(firstCallExpr(t, fi), fi.Src, map[string]bool{"r": true}, map[uintptr]bool{}); routes != nil {
		t.Fatalf("expected nil, got %+v", routes)
	}
}
