//ff:func feature=scan type=test control=sequence
//ff:what TestBuildSpecNode_WithEndpointsCov 테스트
package scanner

import "testing"

func TestBuildSpecNode_WithEndpointsCov(t *testing.T) {
	result := &ScanResult{
		Endpoints: []Endpoint{
			{
				Method: "GET",
				Path:   "/api/users",
				Responses: []Response{
					{Status: "200", Kind: "json", TypeName: "UserResp", Fields: []Field{{Name: "id", JSON: "id", Type: "int"}}},
				},
			},
			{
				Method: "ANY",
				Path:   "/api/health",
			},
		},
	}
	node := buildSpecNode(result)
	if node == nil {
		t.Fatal("expected non-nil")
	}
	// Should have paths and components
	if len(node.Content) < 6 {
		t.Fatalf("expected at least 6 content nodes (openapi+info+paths), got %d", len(node.Content))
	}
}
