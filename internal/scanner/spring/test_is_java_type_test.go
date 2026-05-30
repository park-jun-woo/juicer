//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestIsJavaType 테스트
package spring

import "testing"

func TestIsJavaType(t *testing.T) {
	if !isJavaType("String") || !isJavaType("List<X>") || isJavaType("MyDto") {
		t.Fatal("javaType")
	}
}
