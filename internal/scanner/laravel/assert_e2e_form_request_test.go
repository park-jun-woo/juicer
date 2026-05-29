//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what assertE2EFormRequest — POST /api/users의 FormRequest 본문을 검증한다
package laravel

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func assertE2EFormRequest(t *testing.T, endpoints []scanner.Endpoint) {
	t.Helper()
	ep := findEndpoint(endpoints, "POST", "/api/users")
	if ep == nil {
		return
	}
	if ep.Request == nil || ep.Request.Body == nil {
		t.Error("POST /api/users: expected request body from FormRequest")
		return
	}
	if len(ep.Request.Body.Fields) != 3 {
		t.Errorf("POST /api/users: expected 3 body fields, got %d", len(ep.Request.Body.Fields))
	}
	if ep.Request.Body.TypeName != "StoreUserRequest" {
		t.Errorf("POST /api/users: body type = %q, want StoreUserRequest", ep.Request.Body.TypeName)
	}
}
