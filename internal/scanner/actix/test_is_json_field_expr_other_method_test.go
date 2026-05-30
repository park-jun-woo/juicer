//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestIsJSONFieldExpr_OtherMethod 테스트
package actix

import "testing"

func TestIsJSONFieldExpr_OtherMethod(t *testing.T) {
	src := []byte(`fn f() { x.finish(); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, ".finish")
	field := findChildByType(call, "field_expression")
	if isJSONFieldExpr(field, src) {
		t.Fatal("expected false for non-json field")
	}
}
