//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what 중첩 register prefix가 transitive하게 합성됨을 검증 (/public/v1 + /users)
package fastify

import "testing"

func TestScan_TransitivePrefix(t *testing.T) {
	dir := t.TempDir()

	usersRoutes := `
export default async function(fastify) {
  fastify.get("/list", listUsers);
}
`
	writeFile(t, dir, "routes/users.ts", usersRoutes)

	publicApi := `
import userRoutes from "./users";
export default async function(fastify) {
  fastify.register(userRoutes, { prefix: "/users" });
}
`
	writeFile(t, dir, "routes/public.ts", publicApi)

	appSrc := `
import Fastify from "fastify";
import publicApi from "./routes/public";
const app = Fastify();
app.register(publicApi, { prefix: "/public/v1" });
`
	writeFile(t, dir, "app.ts", appSrc)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}

	found := map[string]bool{}
	for _, ep := range result.Endpoints {
		found[ep.Method+" "+ep.Path] = true
	}
	if !found["GET /public/v1/users/list"] {
		t.Errorf("missing transitive GET /public/v1/users/list, got %v", found)
	}
}
