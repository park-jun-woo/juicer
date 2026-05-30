//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestDefaultStatusForMethod 테스트
package spring

import "testing"

func TestDefaultStatusForMethod(t *testing.T) {
	if defaultStatusForMethod("POST") != "201" || defaultStatusForMethod("GET") != "200" {
		t.Fatal("status")
	}
}
