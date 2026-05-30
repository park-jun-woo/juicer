//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractOpenAPIRoute 테스트
package hono

import "testing"

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
