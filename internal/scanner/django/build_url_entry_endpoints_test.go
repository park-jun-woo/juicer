//ff:func feature=scan type=test control=sequence topic=django
//ff:what buildURLEntryEndpoints — 모듈 URL 맵 전개 후 엔드포인트 생성을 검증
package django

import "testing"

func TestBuildURLEntryEndpoints(t *testing.T) {
	byModule := map[string][]urlEntry{
		"app.urls": {
			{pattern: "health/", viewName: "health"},
		},
	}
	funcViews := []funcViewInfo{{name: "health", methods: []string{"GET"}, file: "v.py"}}
	eps := buildURLEntryEndpoints(byModule, nil, nil, funcViews, map[string]serializerInfo{})
	if len(eps) == 0 {
		t.Fatal("expected endpoints from URL entries")
	}
	if eps[0].Handler != "health" {
		t.Errorf("handler = %q, want health", eps[0].Handler)
	}
}

func TestBuildURLEntryEndpoints_Empty(t *testing.T) {
	eps := buildURLEntryEndpoints(map[string][]urlEntry{}, nil, nil, nil, nil)
	if len(eps) != 0 {
		t.Fatalf("expected no endpoints, got %d", len(eps))
	}
}
