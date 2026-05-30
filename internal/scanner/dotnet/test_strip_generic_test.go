//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestStripGeneric 테스트
package dotnet

import "testing"

func TestStripGeneric(t *testing.T) {
	if stripGeneric("List<X>") != "List" || stripGeneric("X") != "X" {
		t.Fatal("strip")
	}
}
