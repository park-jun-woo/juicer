//ff:func feature=scan type=test control=iteration dimension=1 topic=laravel
//ff:what assertE2EEndpoints — 기대 엔드포인트들이 모두 존재하는지 검증한다
package laravel

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func assertE2EEndpoints(t *testing.T, endpoints []scanner.Endpoint) {
	t.Helper()
	found := endpointKeySet(endpoints)
	expectedEndpoints := []string{
		"GET /api/users",
		"POST /api/users",
		"GET /api/users/{user}",
		"GET /api/posts",
		"POST /api/posts",
		"GET /api/posts/{post}",
		"PUT /api/posts/{post}",
		"DELETE /api/posts/{post}",
		"GET /api/v1/health",
		"GET /api/me",
	}
	for _, expected := range expectedEndpoints {
		if !found[expected] {
			t.Errorf("missing expected endpoint: %s", expected)
		}
	}
}
