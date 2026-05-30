//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestJavaTypeToOpenAPIString 테스트
package quarkus

import "testing"

func TestJavaTypeToOpenAPIString(t *testing.T) {
	if got := javaTypeToOpenAPIString("UUID"); got != "string:uuid" {
		t.Fatalf("got %q", got)
	}
	if got := javaTypeToOpenAPIString("String"); got != "string" {
		t.Fatalf("got %q", got)
	}
}
