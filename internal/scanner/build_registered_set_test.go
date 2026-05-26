//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestBuildRegisteredSet 테스트
package scanner

import "testing"

func TestBuildRegisteredSet(t *testing.T) {
	sr := &ScanResult{
		Endpoints: []Endpoint{
			{Method: "GET", Path: "/api/users"},
			{Method: "ANY", Path: "/api/health"},
		},
	}
	set := buildRegisteredSet(sr)
	if !set["get\t/api/users"] {
		t.Fatal("expected get /api/users in set")
	}
	for _, m := range []string{"get", "post", "put", "patch", "delete"} {
		key := m + "\t/api/health"
		if !set[key] {
			t.Fatalf("expected ANY to register %s /api/health", m)
		}
	}
}
