//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestIsPrimitiveType 테스트
package spring

import "testing"

func TestIsPrimitiveType(t *testing.T) {
	if !isPrimitiveType("int") || isPrimitiveType("MyDto") {
		t.Fatal("primitive")
	}
}
