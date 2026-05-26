//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what buildEndpoint 테스트
package fastapi

import (
	"testing"

	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestBuildEndpoint(t *testing.T) {
	ri := routeInfo{
		method:     "POST",
		path:       "/users",
		handler:    "create_user",
		file:       "main.py",
		line:       10,
		statusCode: 201,
		bodyType:   "UserCreate",
		params:     []scanner.Param{{Name: "id", Type: "integer"}},
	}
	ep := buildEndpoint(ri)
	if ep.Method != "POST" || ep.Path != "/users" || ep.Handler != "create_user" {
		t.Fatalf("basic fields wrong: %+v", ep)
	}
	if ep.Request == nil || ep.Request.Body == nil {
		t.Fatal("expected request body")
	}
	if len(ep.Responses) != 1 || ep.Responses[0].Status != "201" {
		t.Fatalf("unexpected responses: %v", ep.Responses)
	}

	// No body, no response
	ri2 := routeInfo{method: "GET", path: "/health", handler: "health"}
	ep2 := buildEndpoint(ri2)
	if ep2.Request != nil {
		t.Fatal("expected nil request for empty params")
	}
	if len(ep2.Responses) != 0 {
		t.Fatalf("expected 0 responses, got %d", len(ep2.Responses))
	}

	// returnType fallback for response
	ri3 := routeInfo{method: "GET", path: "/me", handler: "get_me", returnType: "UserOut"}
	ep3 := buildEndpoint(ri3)
	if len(ep3.Responses) != 1 {
		t.Fatalf("expected 1 response, got %d", len(ep3.Responses))
	}
}
