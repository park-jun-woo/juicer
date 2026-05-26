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

	// with endpoints that produce paths and schemas
	result2 := &ScanResult{
		Endpoints: []Endpoint{
			{
				Method: "POST",
				Path:   "/api/users",
				Request: &Request{
					Body: &Body{TypeName: "User", Fields: []Field{{Name: "name", JSON: "name", Type: "string"}}},
				},
				Responses: []Response{{Status: "200", Kind: "json", TypeName: "User", Fields: []Field{{Name: "id", JSON: "id", Type: "int"}}}},
			},
			{
				Method: "GET",
				Path:   "/api/users",
			},
		},
	}
	node2 := buildSpecNode(result2)
	if node2 == nil {
		t.Fatal("expected non-nil with endpoints")
	}
	// Check that paths and components/schemas are populated
	if len(node2.Content) <= 6 {
		t.Fatal("expected paths node in content")
	}
}
