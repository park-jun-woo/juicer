//ff:func feature=scan type=test control=iteration dimension=1 topic=actix
//ff:what TestScan_BuilderRoutes — 빌더 패턴 라우트 스캔 테스트
package actix

import "testing"

func TestScan_BuilderRoutes(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/routes.rs", builderRoutesSource)
	writeFile(t, dir, "src/models.rs", builderCreateUserStruct)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	for i, ep := range result.Endpoints {
		t.Logf("ep%d: %s %s handler=%s", i, ep.Method, ep.Path, ep.Handler)
	}
	if len(result.Endpoints) != 3 {
		t.Fatalf("expected 3 endpoints, got %d", len(result.Endpoints))
	}

	// Check endpoints - order may vary but we should have GET /api/users, POST /api/users, GET /api/users/{id}
	found := map[string]bool{}
	for _, ep := range result.Endpoints {
		key := ep.Method + " " + ep.Path
		found[key] = true
	}
	for _, want := range []string{"GET /api/users", "POST /api/users", "GET /api/users/{id}"} {
		if !found[want] {
			t.Errorf("expected endpoint %s not found", want)
		}
	}
}
