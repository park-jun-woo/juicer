//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what 크로스 파일 라우터 테스트: 별도 파일의 router를 import하여 prefix 적용
package express

import "testing"

func TestCrossFileRouter(t *testing.T) {
	dir := t.TempDir()

	routerSrc := `
import express from "express";
const router = express.Router();
router.get("/", list);
router.post("/", create);
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
		t.Errorf("missing GET /api/items")
	}
	if !found["POST /api/items"] {
		t.Errorf("missing POST /api/items")
	}
}
