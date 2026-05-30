//ff:func feature=scan type=test control=sequence topic=hono
//ff:what app.openapi(createRoute({...})) 라우트 추출 테스트
package hono

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestExtractOpenAPIRoute(t *testing.T) {
	src := []byte(`
import { OpenAPIHono, createRoute } from "@hono/zod-openapi"
const app = new OpenAPIHono()
app.openapi(createRoute({ method: "post", path: "/signup" }), signupHandler)
`)
	fi := mustParse(t, src)
	vars := collectHonoVars(fi)
	routes := collectRoutes(fi, vars)
	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d", len(routes))
	}
	r := routes[0]
	if r.Method != "POST" {
		t.Errorf("expected POST, got %s", r.Method)
	}
	if r.Path != "/signup" {
		t.Errorf("expected /signup, got %s", r.Path)
	}
	if r.Handler != "signupHandler" {
		t.Errorf("expected signupHandler, got %s", r.Handler)
	}
}

func openapiCall(t *testing.T, src string) (*fileInfo, *sitter.Node) {
	t.Helper()
	fi := mustParse(t, []byte(src+"\n"))
	// the outer app.openapi(...) call is the first call_expression
	call := findAllByType(fi.Root, "call_expression")[0]
	return fi, call
}

func TestExtractOpenAPIRoute_NoCreateRoute(t *testing.T) {
	// first arg is not createRoute(...) call -> obj nil
	fi, call := openapiCall(t, `app.openapi(routeDef, handler);`)
	if r := extractOpenAPIRoute(call, fi.Src, "app"); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}

func TestExtractOpenAPIRoute_MissingMethodPath(t *testing.T) {
	fi, call := openapiCall(t, `app.openapi(createRoute({ summary: "x" }), handler);`)
	if r := extractOpenAPIRoute(call, fi.Src, "app"); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}

func TestExtractOpenAPIRoute_NoHandlerArg(t *testing.T) {
	// only one arg (createRoute) -> handler stays ""
	fi, call := openapiCall(t, `app.openapi(createRoute({ method: "get", path: "/x" }));`)
	r := extractOpenAPIRoute(call, fi.Src, "app")
	if r == nil || r.Method != "GET" || r.Path != "/x" || r.Handler != "" {
		t.Fatalf("got %+v", r)
	}
	if r.OwnerVar != "app" || r.Line != 1 {
		t.Fatalf("owner/line: %s %d", r.OwnerVar, r.Line)
	}
}

func TestExtractOpenAPIRoute_TwoArgsHandler(t *testing.T) {
	// explicit 2-arg path to exercise the handler extraction branch in isolation
	fi, call := openapiCall(t, `app.openapi(createRoute({ method: "put", path: "/p" }), updateH);`)
	r := extractOpenAPIRoute(call, fi.Src, "app")
	if r == nil || r.Method != "PUT" || r.Handler != "updateH" {
		t.Fatalf("got %+v", r)
	}
}
