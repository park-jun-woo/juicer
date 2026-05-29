//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what assertE2EResourceResponse — GET /api/users/{user}의 Resource 응답을 검증한다
package laravel

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func assertE2EResourceResponse(t *testing.T, endpoints []scanner.Endpoint) {
	t.Helper()
	ep := findEndpoint(endpoints, "GET", "/api/users/{user}")
	if ep == nil {
		return
	}
	if len(ep.Responses) == 0 {
		t.Error("GET /api/users/{user}: expected responses from Resource")
		return
	}
	if ep.Responses[0].TypeName != "UserResource" {
		t.Errorf("GET /api/users/{user}: response type = %q, want UserResource", ep.Responses[0].TypeName)
	}
}
