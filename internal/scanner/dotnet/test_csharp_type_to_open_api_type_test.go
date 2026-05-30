//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestCsharpTypeToOpenAPIType 테스트
package dotnet

import "testing"

func TestCsharpTypeToOpenAPIType(t *testing.T) {
	if csharpTypeToOpenAPIType("int") != "integer" {
		t.Fatal("int")
	}
	if csharpTypeToOpenAPIType("UserDto") != "object" {
		t.Fatal("unknown -> object")
	}
}
