//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestRouteToFieldExpr_To 테스트
package actix

import "testing"

func TestRouteToFieldExpr_To(t *testing.T) {
	src := []byte(`fn f() { web::get().to(h); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, ".to")
	if call == nil {
		t.Fatal("no .to call")
	}
	fe := routeToFieldExpr(call, src)
	if fe == nil || fe.Type() != "field_expression" {
		t.Fatalf("expected field_expression, got %v", fe)
	}
}
