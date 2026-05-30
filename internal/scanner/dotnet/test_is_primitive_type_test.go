//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestIsPrimitiveType 테스트
package dotnet

import "testing"

func TestIsPrimitiveType(t *testing.T) {
	if !isPrimitiveType("int") || isPrimitiveType("UserDto") {
		t.Fatal("primitive")
	}
}
