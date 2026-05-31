//ff:func feature=scan type=test topic=express control=sequence
//ff:what usageRouterCandidate HTTP 메서드+경로 인자일 때만 라우터 변수명 반환 테스트
package express

import "testing"

func TestUsageRouterCandidate(t *testing.T) {
	// valid: router.get('/x', h) -> "router"
	fi := mustParse(t, []byte(`router.get('/x', h);`))
	calls := findAllByType(fi.Root, "call_expression")
	if got := usageRouterCandidate(calls[0], fi.Src); got != "router" {
		t.Errorf("valid: got %q, want router", got)
	}
	// non-HTTP method: config.set(...) -> ""
	fi = mustParse(t, []byte(`config.set('/x', h);`))
	calls = findAllByType(fi.Root, "call_expression")
	if got := usageRouterCandidate(calls[0], fi.Src); got != "" {
		t.Errorf("non-http: got %q", got)
	}
	// HTTP method but non-route first arg: req.get('user-agent') -> ""
	fi = mustParse(t, []byte(`req.get('user-agent');`))
	calls = findAllByType(fi.Root, "call_expression")
	if got := usageRouterCandidate(calls[0], fi.Src); got != "" {
		t.Errorf("non-route arg: got %q", got)
	}
	// blacklisted var: express.get('/x', h) -> ""
	fi = mustParse(t, []byte(`express.get('/x', h);`))
	calls = findAllByType(fi.Root, "call_expression")
	if got := usageRouterCandidate(calls[0], fi.Src); got != "" {
		t.Errorf("blacklist: got %q", got)
	}
}
