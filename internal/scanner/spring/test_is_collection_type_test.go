//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestIsCollectionType 테스트
package spring

import "testing"

func TestIsCollectionType(t *testing.T) {
	if !isCollectionType("List<X>") || isCollectionType("X") {
		t.Fatal("collection")
	}
}
