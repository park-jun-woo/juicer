//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestStripGenericPure 테스트
package spring

import "testing"

func TestStripGenericPure(t *testing.T) {
	if stripGeneric("List<X>") != "List" || stripGeneric("X") != "X" {
		t.Fatal("strip")
	}
}
