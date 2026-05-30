//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestStripGeneric 테스트
package quarkus

import "testing"

func TestStripGeneric(t *testing.T) {
	if stripGeneric("List<String>") != "List" {
		t.Fatal("generic")
	}
	if stripGeneric("String") != "String" {
		t.Fatal("plain")
	}
}
