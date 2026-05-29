//ff:func feature=scan type=test control=sequence topic=express
//ff:what 한 파일의 라우터 둘이 서로 다른 prefix로 마운트되면 각자 prefix를 받음 (Phase097)
package express

import "testing"

// 한 파일에 라우터가 둘이고 서로 다른 prefix로 마운트되면, 각 라우터의
// 라우트는 자기 prefix를 받아야 한다 (파일 단위로 붕괴하면 안 됨).
func TestMultiRouterPerFilePrefix(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "routes/multi.ts", `
import express from "express";
export const aRouter = express.Router();
aRouter.get("/", listA);
export const bRouter = express.Router();
bRouter.get("/", listB);
`)
	writeFile(t, dir, "app.ts", `
import express from "express";
import { aRouter, bRouter } from "./routes/multi";
const app = express();
app.use("/v1/aaa", aRouter);
app.use("/v1/bbb", bRouter);
`)
	writeFile(t, dir, "package.json", `{"dependencies":{"express":"^4.0.0"}}`)
	got := scanSet(t, dir)
	assertHas(t, got, "GET /v1/aaa", "GET /v1/bbb")
}
