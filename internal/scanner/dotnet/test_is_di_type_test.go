//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestIsDIType 테스트
package dotnet

import "testing"

func TestIsDIType(t *testing.T) {
	if !isDIType("ILogger<Foo>") {
		t.Fatal("ILogger")
	}
	if !isDIType("AppDbContext") {
		t.Fatal("DbContext")
	}
	if !isDIType("IUserService") {
		t.Fatal("IService")
	}
	if isDIType("UserDto") {
		t.Fatal("dto not DI")
	}
}
