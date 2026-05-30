//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestIsJavaType 테스트
package quarkus

import "testing"

func TestIsJavaType(t *testing.T) {
	if !isJavaType("String") {
		t.Fatal("String")
	}
	if !isJavaType("List<Foo>") {
		t.Fatal("List")
	}
	if !isJavaType("int[]") {
		t.Fatal("array")
	}
	if isJavaType("MyCustomDto") {
		t.Fatal("custom dto should be false")
	}
}
