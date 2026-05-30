//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestJavaTypeToOpenAPIString 테스트
package spring

import "testing"

func TestJavaTypeToOpenAPIString(t *testing.T) {
	if javaTypeToOpenAPIString("UUID") != "string:uuid" || javaTypeToOpenAPIString("String") != "string" {
		t.Fatal("oa string")
	}
}
