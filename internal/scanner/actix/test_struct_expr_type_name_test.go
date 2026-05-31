//ff:func feature=scan type=test control=sequence topic=actix
//ff:what structExprTypeName struct_expression의 타입명 추출 테스트
package actix

import "testing"

func TestStructExprTypeName(t *testing.T) {
	src := []byte(`fn f() { let x = UserResponse { id: 1 }; }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	se := findAllByType(root, "struct_expression")
	if len(se) == 0 {
		t.Fatal("no struct_expression")
	}
	if got := structExprTypeName(se[0], src); got != "UserResponse" {
		t.Errorf("got %q, want UserResponse", got)
	}
}
