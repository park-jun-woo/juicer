//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what scanPass2 테스트
package fastify

import "testing"

func TestScanPass2(t *testing.T) {
	withRoutes := mustParse(t, []byte(`
const app = Fastify();
app.get("/x", h);
`))
	withRoutes.Path = "/p/a.ts"

	noInstances := mustParse(t, []byte("const x = 1;\n"))
	noInstances.Path = "/p/b.ts"

	ctx := &scanContext{
		parsed: map[string]*fileInfo{
			"/p/a.ts": withRoutes,
			"/p/b.ts": noInstances,
		},
		instances: map[string]map[string]bool{
			"/p/a.ts": {"app": true},
			"/p/b.ts": {}, // empty -> skipped
		},
		prefixMap: map[string][]string{},
		wrappers:  map[string][]wrapperScope{},
		absRoot:   "/p",
	}
	eps := scanPass2(ctx)
	if len(eps) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(eps))
	}
	if eps[0].Path != "/x" {
		t.Errorf("path = %q", eps[0].Path)
	}
}
