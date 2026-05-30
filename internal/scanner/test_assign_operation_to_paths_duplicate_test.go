//ff:func feature=scan type=test control=sequence
//ff:what TestAssignOperationToPaths_Duplicate 테스트
package scanner

import "testing"

func TestAssignOperationToPaths_Duplicate(t *testing.T) {
	paths := map[string]map[string]any{"/users": {"get": map[string]any{"old": true}}}
	op := map[string]any{"new": true}
	ep := Endpoint{Method: "GET", Handler: "list", File: "h.rs", Line: 2}

	assignOperationToPaths(paths, "/users", ep, op)

	got, _ := paths["/users"]["get"].(map[string]any)
	if got["new"] != true {
		t.Fatalf("expected operation overwritten, got %+v", paths["/users"]["get"])
	}
}
