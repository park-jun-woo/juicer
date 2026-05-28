//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what E2E 스캔 테스트: 크로스 파일 플러그인 마운트 + JSON Schema + 경로 파라미터
package fastify

import "testing"

func TestScan_E2E(t *testing.T) {
	dir := t.TempDir()

	userRoutes := `
export default async function(fastify) {
  fastify.get("/", listUsers);
  fastify.get("/:id", getUser);
  fastify.post("/", {
    schema: {
      body: {
        type: "object",
        required: ["name", "email"],
        properties: {
          name: { type: "string" },
          email: { type: "string" }
        }
      },
      response: {
        201: {
          type: "object",
          properties: {
            id: { type: "string" },
            name: { type: "string" }
          }
        }
      }
    }
  }, createUser);
}
`
	writeFile(t, dir, "routes/users.ts", userRoutes)

	appSrc := `
import Fastify from "fastify";
import userRoutes from "./routes/users";
const app = Fastify();
app.register(userRoutes, { prefix: "/api/users" });
app.get("/health", healthCheck);
`
	writeFile(t, dir, "app.ts", appSrc)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) < 4 {
		t.Fatalf("expected at least 4 endpoints, got %d", len(result.Endpoints))
	}

	found := map[string]bool{}
	for _, ep := range result.Endpoints {
		found[ep.Method+" "+ep.Path] = true
	}
	expected := []string{
		"GET /api/users",
		"GET /api/users/{id}",
		"POST /api/users",
		"GET /health",
	}
	for _, e := range expected {
		if !found[e] {
			t.Errorf("missing endpoint %s, got %v", e, found)
		}
	}

	verifyPostEndpoint(t, result)
}
