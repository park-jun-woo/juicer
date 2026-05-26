//ff:func feature=scan type=test control=sequence
//ff:what TestToYAMLNode_BoolCov 테스트
package scanner

import "testing"

func TestToYAMLNode_BoolCov(t *testing.T) {
	node := toYAMLNode(true)
	if node.Value != "true" {
		t.Fatal("expected true")
	}
	node = toYAMLNode(false)
	if node.Value != "false" {
		t.Fatal("expected false")
	}
}
