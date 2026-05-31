//ff:func feature=scan type=test topic=actix control=sequence
//ff:what resolveLetBindingType 블록 내 let 선언에서 변수 타입 해석 테스트
package actix

import "testing"

func TestResolveLetBindingType(t *testing.T) {
	src := []byte(`fn f() { let a = 1; let resp = UserResponse { id: 1 }; }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	block := findAllByType(root, "block")[0]
	if got := resolveLetBindingType("resp", block, src); got != "UserResponse" {
		t.Errorf("got %q", got)
	}
	// unknown var
	if got := resolveLetBindingType("ghost", block, src); got != "" {
		t.Errorf("unknown: got %q", got)
	}
}
