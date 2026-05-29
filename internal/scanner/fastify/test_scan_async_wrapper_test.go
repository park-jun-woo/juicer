//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what async wrapper register({prefix})의 본문 라우트에만 wrapper prefix가 적용됨을 검증
package fastify

import "testing"

func TestScan_AsyncWrapper(t *testing.T) {
	dir := t.TempDir()

	appSrc := `
import Fastify from "fastify";
const app = Fastify();
app.register(async (instance) => {
  instance.get("/items", listItems);
}, { prefix: "/api/v2" });
app.get("/health", healthCheck);
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
	if !found["GET /api/v2/items"] {
		t.Errorf("missing wrapper-scoped GET /api/v2/items, got %v", found)
	}
	if !found["GET /health"] {
		t.Errorf("missing GET /health, got %v", found)
	}
	if found["GET /api/v2/health"] {
		t.Errorf("wrapper prefix leaked to /health, got %v", found)
	}
}
