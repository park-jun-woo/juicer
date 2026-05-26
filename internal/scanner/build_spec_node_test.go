//ff:func feature=scan type=test control=sequence
//ff:what TestBuildSpecNode_Empty 테스트
package scanner

import "testing"

func TestBuildSpecNode_Empty(t *testing.T) {
	result := &ScanResult{Endpoints: nil}
	node := buildSpecNode(result)
	if node == nil {
		t.Fatal("expected non-nil")
	}
}
