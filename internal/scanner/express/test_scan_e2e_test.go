//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what E2E 스캔 테스트: 크로스 파일 라우터 마운트 + 경로 파라미터 + 미들웨어
package express

import "testing"

func TestScan_E2E(t *testing.T) {
	dir := t.TempDir()

	userRoutes := `
import express from "express";
const router = express.Router();
router.get("/", listUsers);
router.get("/:id", getUser);
router.post("/", createUser);
export default router;
`
	writeFile(t, dir, "routes/users.ts", userRoutes)

	appSrc := `
import express from "express";
import usersRouter from "./routes/users";
const app = express();
app.use("/api/users", usersRouter);
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
}
