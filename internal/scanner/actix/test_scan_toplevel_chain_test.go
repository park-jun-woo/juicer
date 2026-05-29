//ff:func feature=scan type=test control=iteration dimension=1 topic=actix
//ff:what TestScan_TopLevelChain — 최상위 .service() 체인/to-only/핸들러명 해석 스캔 테스트
package actix

import "testing"

func TestScan_TopLevelChain(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/routes.rs", topLevelChainSource)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}

	got := map[string]string{}
	for _, ep := range result.Endpoints {
		got[ep.Method+" "+ep.Path] = ep.Handler
	}

	// (a)/(b) top-level chain: all three .service() registrations detected.
	for k, want := range map[string]string{
		"GET /alpha":     "alpha",
		"POST /beta":     "beta",
		"GET /gamma":     "gamma",
		"GET /delta":     "delta",   // App::new() chain
		"POST /epsilon":  "epsilon", // App::new() chain
	} {
		if got[k] != want {
			t.Errorf("%s: handler=%q want %q", k, got[k], want)
		}
	}

	// (c) .to()-only resource expands to ANY methods.
	for _, m := range []string{"GET", "POST", "PUT", "PATCH", "DELETE"} {
		if got[m+" /about"] != "about" {
			t.Errorf("%s /about: handler=%q want about", m, got[m+" /about"])
		}
	}

	// (d) generic_function handler name resolved.
	if got["GET /search"] != "search" {
		t.Errorf("GET /search: handler=%q want search", got["GET /search"])
	}

	// (e) closure route detected with non-empty anonymous handler.
	if h, ok := got["GET /inline"]; !ok || h == "" {
		t.Errorf("GET /inline: missing or empty handler (%q ok=%v)", h, ok)
	}
}
