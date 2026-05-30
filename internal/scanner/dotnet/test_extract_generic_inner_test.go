//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestExtractGenericInner 테스트
package dotnet

import "testing"

func TestExtractGenericInner(t *testing.T) {
	if extractGenericInner("List<UserDto>") != "UserDto" {
		t.Fatal("generic")
	}
	if extractGenericInner("String") != "String" {
		t.Fatal("plain")
	}
}
