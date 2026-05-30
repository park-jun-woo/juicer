//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestIsCollectionType 테스트
package quarkus

import "testing"

func TestIsCollectionType(t *testing.T) {
	if !isCollectionType("List<String>") {
		t.Fatal("List")
	}
	if !isCollectionType("Set<Long>") {
		t.Fatal("Set")
	}
	if isCollectionType("String") {
		t.Fatal("String")
	}
}
