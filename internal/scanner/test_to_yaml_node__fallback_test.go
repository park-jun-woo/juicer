//ff:func feature=scan type=extract control=sequence
//ff:what TestToYAMLNode_Fallback 테스트
package scanner

import "testing"

func TestToYAMLNode_Fallback(t *testing.T) {
	node := toYAMLNode(3.14)
	if node == nil {
		t.Fatal("expected non-nil")
	}
}
