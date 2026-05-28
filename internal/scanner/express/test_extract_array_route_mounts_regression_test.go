//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what 정적 app.use 마운트 회귀 테스트: forEach 추가 후에도 기존 패턴 동작 확인
package express

import "testing"

func TestStaticUseMountStillWorks(t *testing.T) {
	dir := t.TempDir()

	routerSrc := `
import express from "express";
const router = express.Router();
router.get("/", list);
export default router;
`
	writeFile(t, dir, "routes/items.ts", routerSrc)

	appSrc := `
import express from "express";
import itemsRouter from "./routes/items";
const app = express();
app.use("/api/items", itemsRouter);
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
	if !found["GET /api/items"] {
		t.Errorf("missing GET /api/items, got %v", found)
	}
}
