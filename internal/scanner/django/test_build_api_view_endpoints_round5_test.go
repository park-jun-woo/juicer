//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestBuildAPIViewEndpoints_Round5 테스트
package django

import "testing"

func TestBuildAPIViewEndpoints_Round5(t *testing.T) {
	entry := urlEntry{pattern: "ping/<int:pk>/", viewName: "PingView"}
	av := &apiviewInfo{name: "PingView", methods: []string{"GET", "POST"}, file: "views.py", line: 1}
	eps := buildAPIViewEndpoints(entry, av, map[string]serializerInfo{})
	if len(eps) != 2 {
		t.Fatalf("expected 2 endpoints, got %d", len(eps))
	}
	if eps[0].Handler != "PingView.get" {
		t.Errorf("handler: %q", eps[0].Handler)
	}
}
