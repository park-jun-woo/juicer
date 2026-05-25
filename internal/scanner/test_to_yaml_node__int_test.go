//ff:func feature=scan type=extract control=sequence
//ff:what TestToYAMLNode_Int 테스트
package scanner

import "testing"

func TestToYAMLNode_Int(t *testing.T) {
	node := toYAMLNode(42)
	if node.Value != "42" {
		t.Fatal("expected 42")
	}
}
