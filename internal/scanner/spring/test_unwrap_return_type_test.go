//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestUnwrapReturnType 테스트
package spring

import "testing"

func TestUnwrapReturnType(t *testing.T) {
	if _, ok := unwrapReturnType("void"); ok {
		t.Fatal("void")
	}
	if got, ok := unwrapReturnType("ResponseEntity<UserDto>"); ok || got != "UserDto" {
		t.Fatalf("RE single: %q %v", got, ok)
	}
	if got, ok := unwrapReturnType("List<UserDto>"); !ok || got != "UserDto" {
		t.Fatalf("List: %q %v", got, ok)
	}
	if _, ok := unwrapReturnType("ResponseEntity<Void>"); ok {
		t.Fatal("RE Void")
	}
}
