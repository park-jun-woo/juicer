//ff:func feature=scan type=test control=sequence
//ff:what TestAssignOperationToPaths_New 테스트
package scanner

import "testing"

func TestAssignOperationToPaths_New(t *testing.T) {
	paths := map[string]map[string]any{"/users": {}}
	op := map[string]any{"summary": "list"}
	ep := Endpoint{Method: "GET", Handler: "list", File: "h.rs", Line: 1}

	assignOperationToPaths(paths, "/users", ep, op)

	if paths["/users"]["get"] == nil {
		t.Fatal("expected get operation assigned")
	}
}
