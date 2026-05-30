//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractOpenAPIRequest 테스트
package hono

import "testing"

func TestExtractOpenAPIRequest(t *testing.T) {
	src := []byte(`
import { OpenAPIHono, createRoute } from "@hono/zod-openapi"
import { z } from "@hono/zod-openapi"
const app = new OpenAPIHono()
app.openapi(createRoute({
  method: "post",
  path: "/signup",
  request: {
    body: {
      content: {
        "application/json": {
          schema: z.object({ email: z.string(), password: z.string() })
        }
      }
    }
  }
}), signupHandler)
`)
	fi := mustParse(t, src)
	vars := collectHonoVars(fi)
	routes := collectRoutes(fi, vars)
	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d", len(routes))
	}
	r := routes[0]
	if len(r.ZodValidators) != 1 {
		t.Fatalf("expected 1 validator, got %d", len(r.ZodValidators))
	}
	if r.ZodValidators[0].Target != "json" {
		t.Errorf("expected json target, got %s", r.ZodValidators[0].Target)
	}
	if r.ZodValidators[0].SchemaNode == nil {
		t.Error("expected inline schema node")
	}
}
