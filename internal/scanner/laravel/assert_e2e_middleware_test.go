//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what assertE2EMiddleware — GET /api/me의 미들웨어를 검증한다
package laravel

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func assertE2EMiddleware(t *testing.T, endpoints []scanner.Endpoint) {
	t.Helper()
	ep := findEndpoint(endpoints, "GET", "/api/me")
	if ep == nil {
		return
	}
	if len(ep.Middleware) != 1 || ep.Middleware[0] != "auth:sanctum" {
		t.Errorf("GET /api/me: middleware = %v, want [auth:sanctum]", ep.Middleware)
	}
}
