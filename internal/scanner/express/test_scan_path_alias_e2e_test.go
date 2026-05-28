//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what E2E 테스트: tsconfig path alias를 통한 크로스 파일 라우터 마운트
package express

import "testing"

func TestScan_PathAlias_E2E(t *testing.T) {
	dir := t.TempDir()

	tsconfig := `{
  "compilerOptions": {
    "paths": {
      "@/*": ["src/*"]
    }
  }
}`
	writeFile(t, dir, "tsconfig.json", tsconfig)

	userRoutes := `
import express from "express";
const router = express.Router();
router.get("/", listUsers);
router.post("/", createUser);
export default router;
`
	writeFile(t, dir, "src/api/users/router.ts", userRoutes)

	appSrc := `
import express from "express";
import usersRouter from "@/api/users/router";
const app = express();
app.use("/api/users", usersRouter);
app.get("/health", healthCheck);
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
	expected := []string{
		"GET /api/users",
		"POST /api/users",
		"GET /health",
	}
	for _, e := range expected {
		if !found[e] {
			t.Errorf("missing endpoint %s, got %v", e, found)
		}
	}
}
