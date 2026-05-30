//ff:func feature=scan type=test control=iteration dimension=1 topic=actix
//ff:what TestExtractMacroRoutes 테스트
package actix

import "testing"

func TestExtractMacroRoutes(t *testing.T) {
	root, err := parseRust([]byte(macroRoutesSource))
	if err != nil {
		t.Fatal(err)
	}
	fi := &fileInfo{
		absPath: "/x/handlers.rs",
		relPath: "handlers.rs",
		src:     []byte(macroRoutesSource),
		root:    root,
	}

	routes := extractMacroRoutes(fi)
	if len(routes) != 4 {
		t.Fatalf("expected 4 macro routes, got %d: %+v", len(routes), routes)
	}

	got := map[string]string{}
	for _, r := range routes {
		got[r.handler] = r.method + " " + r.path
	}
	want := map[string]string{
		"get_user":    "GET /users/{id}",
		"create_user": "POST /users",
		"update_user": "PUT /users/{id}",
		"delete_user": "DELETE /users/{id}",
	}
	for h, mp := range want {
		if got[h] != mp {
			t.Errorf("route %q = %q, want %q (all: %v)", h, got[h], mp, got)
		}
	}
}
