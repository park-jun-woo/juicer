//ff:func feature=scan type=test control=sequence
//ff:what TestBuildSpecNode_WithEP 테스트
package scanner

import "testing"

func TestBuildSpecNode_WithEP(t *testing.T) {
	result := &ScanResult{
		Endpoints: []Endpoint{
			{Method: "GET", Path: "/api/users", Responses: []Response{{Status: "200", Kind: "json", Fields: []Field{{Name: "id", JSON: "id", Type: "int"}}}}},
		},
	}
	node := buildSpecNode(result)
	if node == nil {
		t.Fatal("expected non-nil")
	}
}
