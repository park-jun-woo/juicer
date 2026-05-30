//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestIsJSONFieldExpr_True 테스트
package actix

import "testing"

func TestIsJSONFieldExpr_True(t *testing.T) {
	src := []byte(`fn f() { HttpResponse::Ok().json(x); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}

	fe := findCallByFuncSuffix(root, src, ".json")
	if fe == nil {
		t.Fatal("no .json call")
	}
	field := findChildByType(fe, "field_expression")
	if !isJSONFieldExpr(field, src) {
		t.Fatal("expected isJSONFieldExpr true")
	}
}
