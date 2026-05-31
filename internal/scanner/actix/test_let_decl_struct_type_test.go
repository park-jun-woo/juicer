//ff:func feature=scan type=test topic=actix control=sequence
//ff:what letDeclStructType let 선언 변수명 일치 시 struct 타입명 반환 테스트
package actix

import "testing"

func TestLetDeclStructType(t *testing.T) {
	src := []byte(`fn f() { let resp = UserResponse { id: 1 }; }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	decl := findAllByType(root, "let_declaration")[0]
	if got := letDeclStructType(decl, "resp", src); got != "UserResponse" {
		t.Errorf("match: got %q", got)
	}
	// name mismatch
	if got := letDeclStructType(decl, "other", src); got != "" {
		t.Errorf("mismatch: got %q", got)
	}
	// non-struct value
	src2 := []byte(`fn f() { let n = 5; }`)
	root2, _ := parseRust(src2)
	decl2 := findAllByType(root2, "let_declaration")[0]
	if got := letDeclStructType(decl2, "n", src2); got != "" {
		t.Errorf("non-struct: got %q", got)
	}
}
