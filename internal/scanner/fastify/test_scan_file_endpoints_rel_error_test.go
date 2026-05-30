//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestScanFileEndpoints_RelError 테스트
package fastify

import "testing"

func TestScanFileEndpoints_RelError(t *testing.T) {

	src := `
const app = Fastify();
app.get("/x", h);
`
	fi := mustParse(t, []byte(src))
	path := "/elsewhere/app.ts"
	fi.Path = path
	ctx := &scanContext{
		parsed:    map[string]*fileInfo{path: fi},
		prefixMap: map[string][]string{},
		wrappers:  map[string][]wrapperScope{},
		absRoot:   "relative-not-absolute",
	}
	eps := scanFileEndpoints(ctx, path, fi, map[string]bool{"app": true})
	if len(eps) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(eps))
	}
	if eps[0].File != path {
		t.Errorf("expected fallback to full path %q, got %q", path, eps[0].File)
	}
}
