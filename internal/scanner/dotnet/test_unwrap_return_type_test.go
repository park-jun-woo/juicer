//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestUnwrapReturnType 테스트
package dotnet

import "testing"

func TestUnwrapReturnType(t *testing.T) {
	if got, ok := unwrapReturnType("ActionResult<UserDto>"); ok || got != "UserDto" {
		t.Fatalf("ActionResult: %q %v", got, ok)
	}
	if got, ok := unwrapReturnType("Task<List<UserDto>>"); !ok || got != "UserDto" {
		t.Fatalf("Task<List>: %q %v", got, ok)
	}
	if got, ok := unwrapReturnType("UserDto[]"); !ok || got != "UserDto" {
		t.Fatalf("array: %q %v", got, ok)
	}
}
