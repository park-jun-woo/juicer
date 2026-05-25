//ff:func feature=scan type=extract control=sequence
//ff:what TestBuildSpecNode_WithEndpoints 테스트
package scanner

import "testing"

func TestBuildSpecNode_WithEndpoints(t *testing.T) {
	result := &ScanResult{
		Endpoints: []Endpoint{
			{Method: "GET", Path: "/api/health", Responses: []Response{{Status: "200"}}},
		},
	}
	node := buildSpecNode(result)
	if node == nil {
		t.Fatal("expected non-nil")
	}
}
