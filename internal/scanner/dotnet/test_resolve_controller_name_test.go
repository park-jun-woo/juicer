//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestResolveControllerName 테스트
package dotnet

import "testing"

func TestResolveControllerName(t *testing.T) {
	if resolveControllerName("UsersController") != "users" {
		t.Fatal("controller name")
	}
}
