//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestIntPtr 테스트
package dotnet

import "testing"

func TestIntPtr(t *testing.T) {
	if p := intPtr(8); p == nil || *p != 8 {
		t.Fatalf("got %v", p)
	}
}
