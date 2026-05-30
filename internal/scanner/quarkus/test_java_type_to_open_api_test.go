//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestJavaTypeToOpenAPI 테스트
package quarkus

import "testing"

func TestJavaTypeToOpenAPI(t *testing.T) {
	if javaTypeToOpenAPI("String").Type != "string" {
		t.Fatal("string")
	}
	if javaTypeToOpenAPI("Long").Format != "int64" {
		t.Fatal("long")
	}
	if javaTypeToOpenAPI("UUID").Format != "uuid" {
		t.Fatal("uuid")
	}
}
