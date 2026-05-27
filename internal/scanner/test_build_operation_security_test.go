//ff:func feature=scan type=test control=sequence
//ff:what TestBuildOperation_Security 테스트
package scanner

import "testing"

func TestBuildOperation_Security(t *testing.T) {
	schemas := map[string]any{}

	// auth_required endpoint should have security
	ep := Endpoint{
		Method:    "GET",
		Path:      "/api/users",
		AuthLevel: "auth_required",
	}
	op := buildOperation(ep, schemas)
	sec, ok := op["security"]
	if !ok {
		t.Fatal("expected security field for auth_required endpoint")
	}
	secList, ok := sec.([]any)
	if !ok || len(secList) != 1 {
		t.Fatal("expected security to be a list with one element")
	}
	secMap, ok := secList[0].(map[string]any)
	if !ok {
		t.Fatal("expected security element to be a map")
	}
	if _, ok := secMap["bearerAuth"]; !ok {
		t.Fatal("expected bearerAuth in security element")
	}

	// public endpoint should not have security
	ep2 := Endpoint{
		Method:    "GET",
		Path:      "/api/health",
		AuthLevel: "public",
	}
	op2 := buildOperation(ep2, schemas)
	if _, ok := op2["security"]; ok {
		t.Fatal("expected no security field for public endpoint")
	}

	// middleware-based auth
	ep3 := Endpoint{
		Method:     "GET",
		Path:       "/api/profile",
		Middleware: []string{"JwtAuthGuard"},
	}
	op3 := buildOperation(ep3, schemas)
	if _, ok := op3["security"]; !ok {
		t.Fatal("expected security field for middleware auth endpoint")
	}

	// no auth info at all
	ep4 := Endpoint{
		Method: "GET",
		Path:   "/api/public",
	}
	op4 := buildOperation(ep4, schemas)
	if _, ok := op4["security"]; ok {
		t.Fatal("expected no security field for endpoint without auth")
	}
}
