//ff:func feature=scan type=extract control=sequence
//ff:what TestToYAMLNode_Bool 테스트
package scanner

import "testing"

func TestToYAMLNode_Bool(t *testing.T) {
	node := toYAMLNode(true)
	if node.Value != "true" {
		t.Fatal("expected true")
	}
}
