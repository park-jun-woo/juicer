//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestUnwrapReturnType 테스트
package quarkus

import "testing"

func TestUnwrapReturnType(t *testing.T) {
	if _, ok := unwrapReturnType("void"); ok {
		t.Fatal("void")
	}
	if _, ok := unwrapReturnType("Response"); ok {
		t.Fatal("Response")
	}

	if got, ok := unwrapReturnType("Uni<UserDto>"); ok || got != "UserDto" {
		t.Fatalf("Uni: %q %v", got, ok)
	}

	if got, ok := unwrapReturnType("List<UserDto>"); !ok || got != "UserDto" {
		t.Fatalf("List: %q %v", got, ok)
	}

	if got, ok := unwrapReturnType("Multi<UserDto>"); !ok || got != "UserDto" {
		t.Fatalf("Multi: %q %v", got, ok)
	}
}
