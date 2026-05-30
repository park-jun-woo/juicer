//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestDefaultStatusForMethod 테스트
package dotnet

import "testing"

func TestDefaultStatusForMethod(t *testing.T) {
	if defaultStatusForMethod("POST") != "201" || defaultStatusForMethod("GET") != "200" {
		t.Fatal("status")
	}
}
