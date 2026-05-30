//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestCsharpTypeToOpenAPI 테스트
package dotnet

import "testing"

func TestCsharpTypeToOpenAPI(t *testing.T) {
	if csharpTypeToOpenAPI("string").Type != "string" {
		t.Fatal("string")
	}
	if csharpTypeToOpenAPI("Guid").Format != "uuid" {
		t.Fatal("guid")
	}
	arr := csharpTypeToOpenAPI("List<UserDto>")
	if arr.Type != "array" || arr.Items != "UserDto" {
		t.Fatalf("array: %+v", arr)
	}
	if csharpTypeToOpenAPI("int?").Type != "integer" {
		t.Fatal("nullable int")
	}
}
