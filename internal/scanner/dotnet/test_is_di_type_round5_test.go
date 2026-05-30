//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestIsDIType_Round5 테스트
package dotnet

import "testing"

func TestIsDIType_Round5(t *testing.T) {
	if !isDIType("ILogger<Foo>") {
		t.Error("ILogger should be DI")
	}
	if !isDIType("AppDbContext") {
		t.Error("DbContext should be DI")
	}
	if !isDIType("IUserService") {
		t.Error("IUserService should be DI")
	}
	if !isDIType("IOrderRepository") {
		t.Error("IOrderRepository should be DI")
	}
	if isDIType("IUserThing") {
		t.Error("IUserThing should not be DI")
	}
	if isDIType("string") {
		t.Error("string should not be DI")
	}
	if isDIType("I") {
		t.Error("single I should not be DI")
	}
	if isDIType("iLower") {
		t.Error("lowercase second char should not be DI")
	}
}
