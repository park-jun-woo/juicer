//ff:func feature=scan type=test topic=actix control=iteration dimension=1
//ff:what findJSONFieldExpr scoped_identifier에서 .json field_expression 탐색 테스트
package actix

import "testing"

func TestFindJSONFieldExpr(t *testing.T) {
	src := []byte(`fn h() { HttpResponse::Ok().json(x) }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	scoped := findAllByType(root, "scoped_identifier")
	found := false
	for _, sid := range scoped {
		if fe := findJSONFieldExpr(sid, src); fe != nil && fe.Type() == "field_expression" {
			found = true
		}
	}
	if !found {
		t.Error(".json field_expression not found")
	}

	// a scoped identifier not in a json chain -> nil
	src2 := []byte(`fn h() { std::mem::drop(x) }`)
	root2, _ := parseRust(src2)
	for _, sid := range findAllByType(root2, "scoped_identifier") {
		if findJSONFieldExpr(sid, src2) != nil {
			t.Error("non-json should be nil")
		}
	}
}
