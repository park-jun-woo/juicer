//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what 동일 플러그인을 두 prefix로 등록하면 prefix별로 별도 emit되고 연결되지 않음을 검증
package fastify

import "testing"

func TestScan_MultiRegister(t *testing.T) {
	dir := t.TempDir()

	bootRoutes := `
export default async function(fastify) {
  fastify.get("/status", statusHandler);
}
`
	writeFile(t, dir, "routes/boot.ts", bootRoutes)

	appSrc := `
import Fastify from "fastify";
import bootRoutes from "./routes/boot";
const app = Fastify();
app.register(bootRoutes, { prefix: "/boot" });
app.register(bootRoutes, { prefix: "/new_boot" });
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
	if !found["GET /boot/status"] {
		t.Errorf("missing GET /boot/status, got %v", found)
	}
	if !found["GET /new_boot/status"] {
		t.Errorf("missing GET /new_boot/status, got %v", found)
	}
	if found["GET /boot/new_boot/status"] {
		t.Errorf("must not join prefixes: GET /boot/new_boot/status present, got %v", found)
	}
}
