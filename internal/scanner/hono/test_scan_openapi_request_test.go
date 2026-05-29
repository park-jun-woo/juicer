//ff:func feature=scan type=test control=sequence topic=hono
//ff:what E2E: app.openapi(createRoute) request.body가 endpoint Request body로 추출되는지 테스트
package hono

import "testing"

func TestScan_OpenAPIRequest(t *testing.T) {
	dir := t.TempDir()

	appSrc := `
import { OpenAPIHono, createRoute, z } from "@hono/zod-openapi"

const app = new OpenAPIHono()

app.openapi(
  createRoute({
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
  }),
  (c) => c.json({ ok: true })
)
`
	writeFile(t, dir, "app.ts", appSrc)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(result.Endpoints))
	}
	ep := result.Endpoints[0]
	if ep.Method != "POST" || ep.Path != "/signup" {
		t.Fatalf("expected POST /signup, got %s %s", ep.Method, ep.Path)
	}
	if ep.Request == nil || ep.Request.Body == nil {
		t.Fatal("expected request body")
	}
	if len(ep.Request.Body.Fields) != 2 {
		t.Errorf("expected 2 body fields, got %d", len(ep.Request.Body.Fields))
	}
}
