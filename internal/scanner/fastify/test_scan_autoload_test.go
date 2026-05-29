//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what @fastify/autoload 디렉터리 구조가 prefix로 변환됨을 검증
package fastify

import "testing"

func TestScan_Autoload(t *testing.T) {
	dir := t.TempDir()

	authRoutes := `
export default async function(fastify) {
  fastify.post("/login", login);
}
`
	writeFile(t, dir, "src/routes/api/auth/index.ts", authRoutes)

	tasksRoutes := `
export default async function(fastify) {
  fastify.get("/", listTasks);
}
`
	writeFile(t, dir, "src/routes/api/tasks/index.ts", tasksRoutes)

	appSrc := `
import Fastify from "fastify";
import autoload from "@fastify/autoload";
import { join } from "path";
const app = Fastify();
app.register(autoload, {
  dir: join(__dirname, "routes"),
  options: { prefix: "/api" }
});
`
	writeFile(t, dir, "src/app.ts", appSrc)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}

	found := map[string]bool{}
	for _, ep := range result.Endpoints {
		found[ep.Method+" "+ep.Path] = true
	}
	if !found["POST /api/api/auth/login"] {
		t.Errorf("missing autoload POST /api/api/auth/login, got %v", found)
	}
	if !found["GET /api/api/tasks"] {
		t.Errorf("missing autoload GET /api/api/tasks, got %v", found)
	}
}
