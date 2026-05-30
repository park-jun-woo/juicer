//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestCompositeLitTypeName_Variants 테스트
package echo

import "testing"

func TestCompositeLitTypeName_Variants(t *testing.T) {
	d, b := compositeLitTypeName("UserDto{}")
	if d != "UserDto" || b != "UserDto" {
		t.Fatalf("got %q %q", d, b)
	}
	d2, b2 := compositeLitTypeName("[]UserDto{}")
	if d2 != "[]UserDto" || b2 != "UserDto" {
		t.Fatalf("slice: %q %q", d2, b2)
	}
	if d3, _ := compositeLitTypeName("notcomposite"); d3 != "" {
		t.Fatalf("non-composite: %q", d3)
	}
	if d4, _ := compositeLitTypeName("pkg.Type{}"); d4 != "" {
		t.Fatalf("qualified should be rejected: %q", d4)
	}
}
