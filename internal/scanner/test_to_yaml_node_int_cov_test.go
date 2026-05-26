//ff:func feature=scan type=test control=sequence
//ff:what TestToYAMLNode_IntCov 테스트
package scanner

import "testing"

func TestToYAMLNode_IntCov(t *testing.T) {
	node := toYAMLNode(42)
	if node.Value != "42" {
		t.Fatal("expected 42")
	}
}
