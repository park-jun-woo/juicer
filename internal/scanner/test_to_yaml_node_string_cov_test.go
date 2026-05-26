//ff:func feature=scan type=test control=sequence
//ff:what TestToYAMLNode_StringCov 테스트
package scanner

import "testing"

func TestToYAMLNode_StringCov(t *testing.T) {
	node := toYAMLNode("hello")
	if node.Value != "hello" {
		t.Fatal("expected hello")
	}
}
