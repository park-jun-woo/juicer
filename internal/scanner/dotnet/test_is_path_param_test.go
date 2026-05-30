//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestIsPathParam 테스트
package dotnet

import "testing"

func TestIsPathParam(t *testing.T) {
	if !isPathParam("id", "/users/{id}") || isPathParam("id", "/users") {
		t.Fatal("path param")
	}
}
