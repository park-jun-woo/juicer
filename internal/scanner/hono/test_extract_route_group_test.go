//ff:func feature=scan type=test topic=hono control=sequence
//ff:what extractRouteGroup app.route('/api', sub) 그룹 추출 및 비대상 무시 테스트
package hono

import "testing"

func TestExtractRouteGroup(t *testing.T) {
	honoVars := map[string]bool{"app": true}

	fi := mustParse(t, []byte(`app.route('/api', apiRouter)`))
	call := findAllByType(fi.Root, "call_expression")[0]
	rg := extractRouteGroup(call, fi.Src, honoVars)
	if rg == nil {
		t.Fatalf("route group not extracted")
	}

	// not .route
	fi2 := mustParse(t, []byte(`app.get('/x', h)`))
	c2 := findAllByType(fi2.Root, "call_expression")[0]
	if extractRouteGroup(c2, fi2.Src, honoVars) != nil {
		t.Error(".get is not a route group")
	}

	// not a hono var
	fi3 := mustParse(t, []byte(`other.route('/x', r)`))
	c3 := findAllByType(fi3.Root, "call_expression")[0]
	if extractRouteGroup(c3, fi3.Src, honoVars) != nil {
		t.Error("non-hono var should be nil")
	}
}
