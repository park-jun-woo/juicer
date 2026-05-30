//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestDefaultStatusForMethod 테스트
package quarkus

import "testing"

func TestDefaultStatusForMethod(t *testing.T) {
	if defaultStatusForMethod("POST") != "201" {
		t.Fatal("POST")
	}
	if defaultStatusForMethod("GET") != "200" {
		t.Fatal("GET")
	}
}
