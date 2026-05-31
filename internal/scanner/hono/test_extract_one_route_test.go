//ff:func feature=scan type=test topic=hono control=sequence
//ff:what extractOneRoute app.get('/x', h) 라우트 추출 및 비-hono/비메서드 무시 테스트
package hono

import "testing"

func TestExtractOneRoute(t *testing.T) {
	honoVars := map[string]bool{"app": true}

	// app.get('/users', handler)
	fi := mustParse(t, []byte(`app.get('/users', handler)`))
	call := findAllByType(fi.Root, "call_expression")[0]
	ri := extractOneRoute(call, fi.Src, honoVars)
	if ri == nil || ri.Method != "GET" || ri.Path != "/users" {
		t.Fatalf("route: %+v", ri)
	}

	// not a hono var
	fi2 := mustParse(t, []byte(`other.get('/x', h)`))
	c2 := findAllByType(fi2.Root, "call_expression")[0]
	if extractOneRoute(c2, fi2.Src, honoVars) != nil {
		t.Error("non-hono var should be nil")
	}

	// not an http method
	fi3 := mustParse(t, []byte(`app.use('/x', mw)`))
	c3 := findAllByType(fi3.Root, "call_expression")[0]
	if r := extractOneRoute(c3, fi3.Src, honoVars); r != nil && r.Method != "" {
		t.Errorf("use is not a route method: %+v", r)
	}
}
