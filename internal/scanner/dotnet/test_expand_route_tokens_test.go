//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestExpandRouteTokens 테스트
package dotnet

import "testing"

func TestExpandRouteTokens(t *testing.T) {
	if got := expandRouteTokens("api/[controller]", "UsersController", ""); got != "api/users" {
		t.Fatalf("got %q", got)
	}
}
