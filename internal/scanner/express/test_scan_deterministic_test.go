//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what 결정성 회귀 테스트: 동일 입력을 여러 번 스캔해도 endpoint 시퀀스가 바이트 단위로 동일함을 보장 (Phase096)
package express

import (
	"fmt"
	"testing"
)

// 여러 파일 + 다단계 라우터 마운트. 과거에는 scanPass2가 map(ctx.parsed)을
// 순회해 endpoint 순서가 매 실행 무작위였다. (Phase096)
func TestScan_Deterministic(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, dir, "routes/orgs.ts", `
import express from "express";
const router = express.Router();
router.get("/", listOrgs);
router.post("/", createOrg);
export default router;
`)
	writeFile(t, dir, "routes/licenses.ts", `
import express from "express";
const router = express.Router();
router.get("/", listLicenses);
router.get("/:id", getLicense);
export default router;
`)
	writeFile(t, dir, "routes/activity.ts", `
import express from "express";
const router = express.Router();
router.get("/", listActivityLogs);
export default router;
`)
	writeFile(t, dir, "routes/v1.ts", `
import express from "express";
import orgs from "./orgs";
import licenses from "./licenses";
import activity from "./activity";
const router = express.Router();
router.use("/orgs", orgs);
router.use("/licenses", licenses);
router.use("/activity", activity);
export default router;
`)
	writeFile(t, dir, "app.ts", `
import express from "express";
import v1 from "./routes/v1";
const app = express();
app.use("/v1", v1);
app.get("/health", healthCheck);
`)

	const runs = 30
	var first string
	for i := 0; i < runs; i++ {
		result, err := Scan(dir)
		if err != nil {
			t.Fatalf("run %d: Scan error: %v", i, err)
		}
		var sig string
		for _, ep := range result.Endpoints {
			sig += fmt.Sprintf("%s %s | %s | %s:%d\n", ep.Method, ep.Path, ep.Handler, ep.File, ep.Line)
		}
		if i == 0 {
			first = sig
			continue
		}
		if sig != first {
			t.Fatalf("non-deterministic output on run %d:\n--- run 0 ---\n%s\n--- run %d ---\n%s", i, first, i, sig)
		}
	}
}
