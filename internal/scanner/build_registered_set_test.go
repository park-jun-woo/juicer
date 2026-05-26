//ff:func feature=scan type=test control=sequence
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
	if !set["get\t/api/health"] {
		t.Fatal("expected ANY to map to get")
	}
}
