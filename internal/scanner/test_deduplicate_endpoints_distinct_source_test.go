//ff:func feature=scan type=test control=sequence
//ff:what TestDeduplicateEndpoints_DistinctSource: 출처 파일이 다른 동일 (method,path)는 모두 보존 (BUG-009 / Phase138-A)
package scanner

import "testing"

func TestDeduplicateEndpoints_DistinctSource(t *testing.T) {
	// 같은 (GET, /posts)지만 출처 파일이 다르면 둘 다 보존된다 (소실 안전망).
	eps := []Endpoint{
		{Method: "GET", Path: "/posts", File: "admin/routes.js"},
		{Method: "GET", Path: "/posts", File: "content/routes.js"},
	}
	got := DeduplicateEndpoints(eps)
	if len(got) != 2 {
		t.Fatalf("expected 2 preserved (distinct source), got %d", len(got))
	}

	// 같은 출처 파일 내 동일 (method,path)는 종전대로 합쳐진다 (richer 우선).
	eps2 := []Endpoint{
		{Method: "GET", Path: "/posts", File: "content/routes.js"},
		{Method: "GET", Path: "/posts", File: "content/routes.js", Responses: []Response{{Status: "200", Kind: "json", Fields: []Field{{Name: "id"}}}}},
	}
	got2 := DeduplicateEndpoints(eps2)
	if len(got2) != 1 {
		t.Fatalf("expected 1 (same source dedup), got %d", len(got2))
	}
	if len(got2[0].Responses) == 0 {
		t.Fatal("expected richer endpoint to win within same source")
	}
}
