//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestIsPrimitiveType 테스트
package quarkus

import "testing"

func TestIsPrimitiveType(t *testing.T) {
	if !isPrimitiveType("int") {
		t.Fatal("int")
	}
	if isPrimitiveType("MyDto") {
		t.Fatal("dto")
	}
}
