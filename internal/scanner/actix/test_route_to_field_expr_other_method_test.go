//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestRouteToFieldExpr_OtherMethod 테스트
package actix

import "testing"

func TestRouteToFieldExpr_OtherMethod(t *testing.T) {

	src := []byte(`fn f() { x.route(a); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, ".route")
	if fe := routeToFieldExpr(call, src); fe != nil {
		t.Fatalf("expected nil for non-to method, got %v", fe)
	}
}
