//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestRouteToFieldExpr_NoFieldExpr 테스트
package actix

import "testing"

func TestRouteToFieldExpr_NoFieldExpr(t *testing.T) {

	src := []byte(`fn f() { web::get(); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, "::get")
	if fe := routeToFieldExpr(call, src); fe != nil {
		t.Fatalf("expected nil when no field_expression, got %v", fe)
	}
}
