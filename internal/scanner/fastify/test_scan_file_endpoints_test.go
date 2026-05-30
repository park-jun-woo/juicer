//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestScanFileEndpoints 테스트
package fastify

import "testing"

func TestScanFileEndpoints(t *testing.T) {
	src := `
import Fastify from "fastify";
const app = Fastify();
app.get("/health", health);
app.post("/users", createUser);
`
	fi := mustParse(t, []byte(src))
	path := "/proj/app.ts"
	fi.Path = path
	ctx := &scanContext{
		parsed:    map[string]*fileInfo{path: fi},
		instances: map[string]map[string]bool{path: {"app": true}},
		prefixMap: map[string][]string{path: {"/api"}},
		wrappers:  map[string][]wrapperScope{},
		absRoot:   "/proj",
	}
	eps := scanFileEndpoints(ctx, path, fi, map[string]bool{"app": true})
	if len(eps) != 2 {
		t.Fatalf("expected 2 endpoints, got %d", len(eps))
	}
	for _, ep := range eps {
		if ep.File != "app.ts" {
			t.Errorf("expected relative file app.ts, got %q", ep.File)
		}
		if ep.Path != "/api/health" && ep.Path != "/api/users" {
			t.Errorf("unexpected path %q", ep.Path)
		}
	}
}
