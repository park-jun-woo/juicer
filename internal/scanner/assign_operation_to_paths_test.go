//ff:func feature=scan type=test control=sequence
//ff:what assignOperationToPaths — operation 할당 및 중복 경고 분기를 검증
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

func TestAssignOperationToPaths_Duplicate(t *testing.T) {
	paths := map[string]map[string]any{"/users": {"get": map[string]any{"old": true}}}
	op := map[string]any{"new": true}
	ep := Endpoint{Method: "GET", Handler: "list", File: "h.rs", Line: 2}

	// Duplicate get -> warning path, then overwrite.
	assignOperationToPaths(paths, "/users", ep, op)

	got, _ := paths["/users"]["get"].(map[string]any)
	if got["new"] != true {
		t.Fatalf("expected operation overwritten, got %+v", paths["/users"]["get"])
	}
}

func TestAssignOperationToPaths_Any(t *testing.T) {
	paths := map[string]map[string]any{"/x": {}}
	ep := Endpoint{Method: "any", Handler: "h"}
	assignOperationToPaths(paths, "/x", ep, map[string]any{"k": 1})

	for _, m := range []string{"get", "post", "put", "patch", "delete"} {
		if paths["/x"][m] == nil {
			t.Errorf("expected method %q assigned for any", m)
		}
	}
}
