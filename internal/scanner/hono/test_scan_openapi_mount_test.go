//ff:func feature=scan type=test control=iteration dimension=1 topic=hono
//ff:what E2E: OpenAPIHono + createRoute + .tsx + cross-file app.route 마운트 prefix 합성
package hono

import "testing"

func TestScan_OpenAPIMount(t *testing.T) {
	dir := t.TempDir()

	indexSrc := `
import { OpenAPIHono } from "@hono/zod-openapi"
import authRoutes from "./auth/routes"

const app = new OpenAPIHono()

app.route("/auth", authRoutes)
app.get("/", (c) => c.text("ok"))
app.get("/ui", (c) => c.text("ui"))
`
	authSrc := `
import { OpenAPIHono, createRoute } from "@hono/zod-openapi"

const app = new OpenAPIHono()

app.openapi(createRoute({ method: "post", path: "/signup" }), signup)
app.openapi(createRoute({ method: "post", path: "/login" }), login)
app.get("/logout", (c) => c.text("bye"))

export default app
`
	writeFile(t, dir, "index.ts", indexSrc)
	writeFile(t, dir, "auth/routes.tsx", authSrc)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}

	found := map[string]bool{}
	for _, ep := range result.Endpoints {
		found[ep.Method+" "+ep.Path] = true
	}
	expected := []string{
		"GET /",
		"GET /ui",
		"POST /auth/signup",
		"POST /auth/login",
		"GET /auth/logout",
	}
	for _, e := range expected {
		if !found[e] {
			t.Errorf("missing endpoint %s, got %v", e, found)
		}
	}
}
